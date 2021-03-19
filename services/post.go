package services

import (
	"fmt"
	"github.com/OhMinsSup/story-server/app"
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/ent"
	postEnt "github.com/OhMinsSup/story-server/ent/post"
	tagEnt "github.com/OhMinsSup/story-server/ent/tag"
	userEnt "github.com/OhMinsSup/story-server/ent/user"
	"github.com/OhMinsSup/story-server/libs"
	"github.com/SKAhack/go-shortid"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"log"
	"net/http"
	"time"
)

// 공용 스토어
type PostStore struct {
	tx  *ent.Tx
	ctx *gin.Context
}

func NewPostStore(ctx *gin.Context, tx *ent.Tx) *PostStore {
	return &PostStore{
		tx:  tx,
		ctx: ctx,
	}
}

// UpdatePostService - 포스트 수정 서비스
func UpdatePostService(body dto.UpdatePostDTO, ctx *gin.Context) (*app.ResponseException, error) {
	client := ctx.MustGet("client").(*ent.Client)

	tx, err := client.Tx(ctx)
	if err != nil {
		return app.TransactionsErrorResponse(err.Error(), nil), nil
	}

	store := NewPostStore(ctx, tx)
	log.Println(store)

	userId, err := uuid.Parse(ctx.MustGet("id").(string))
	if err != nil {
		return app.UnAuthorizedErrorResponse("INVALID_USER_ID_UUID", nil), nil
	}

	postId, err := uuid.Parse(body.PostID)
	if err != nil {
		return app.BadRequestErrorResponse("INVALID_POST_ID_UUID", nil), nil
	}

	post, err := store.
		tx.
		Post.
		Query().
		WithTags().
		WithUser().
		Where(postEnt.IDEQ(postId)).
		Only(store.ctx)

	if err != nil {
		return app.NotFoundErrorResponse("Post is not found", nil), nil
	}

	if post.Edges.User.ID != userId {
		return app.ForbiddenErrorResponse("This post is not yours", nil), nil
	}

	updateQuery := post.
		Update().
		SetTitle(body.Title).
		SetBody(body.Body).
		SetIsPrivate(body.IsPrivate).
		SetIsTemp(body.IsTemp).
		SetIsMarkdown(body.IsMarkdown).
		SetMeta(body.Meta).
		SetThumbnail(body.Thumbnail)

	if post.IsTemp && !body.IsTemp {
		now := time.Now()
		updateQuery.SetReleasedAt(now)
	}

	processedUrlSlug := body.UrlSlug
	urlSlugDuplicate, err := tx.Post.Query().Where(
		postEnt.And(
			postEnt.HasUserWith(
				userEnt.IDEQ(userId),
			),
			postEnt.URLSlug(body.UrlSlug),
		),
	).First(ctx)

	if !ent.IsNotFound(err) && urlSlugDuplicate.ID != post.ID {
		processedUrlSlug = generateUrlSlug(body.Title)
	}

	if processedUrlSlug == "" {
		processedUrlSlug = generateUrlSlug(body.Title)
	}

	updateQuery.SetURLSlug(processedUrlSlug)

	updatePost, err := updateQuery.Save(ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return app.TransactionsErrorResponse(rerr.Error(), nil), nil
		}
		return app.InteralServerErrorResponse(err.Error(), nil), nil
	}
	log.Println("update post", updatePost)
	log.Println("tags", body.Tags)
	tagObj, err := syncPostTags(post.ID, body.Tags, store)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return app.TransactionsErrorResponse(rerr.Error(), nil), nil
		}
		return app.InteralServerErrorResponse(err.Error(), nil), nil
	}

	// 태그 객체는 nil 일수도 있다. 그렇기 때문에 nil 체크를 한다.
	if tagObj != nil {
		missings := tagObj["missing"].([]uuid.UUID)
		addings := tagObj["adding"].([]uuid.UUID)

		log.Println("missings", missings)
		log.Println("adding", addings)

		if len(missings) > 0 {
			_, err := post.
				Update().
				RemoveTagIDs(missings...).
				Save(ctx)

			if err != nil {
				if rerr := tx.Rollback(); rerr != nil {
					return app.TransactionsErrorResponse(rerr.Error(), nil), nil
				}
				return app.InteralServerErrorResponse(err.Error(), nil), nil
			}
		}

		if len(addings) > 0 {
			_, err := post.
				Update().
				AddTagIDs(addings...).
				Save(ctx)

			if err != nil {
				if rerr := tx.Rollback(); rerr != nil {
					return app.TransactionsErrorResponse(rerr.Error(), nil), nil
				}
				return app.InteralServerErrorResponse(err.Error(), nil), nil
			}
		}
	}

	return &app.ResponseException{
		Code:          http.StatusOK,
		ResultCode:    0,
		Message:       "",
		ResultMessage: "",
		Data: libs.JSON{
			"id": postId,
		},
	}, tx.Commit()
}

// WritePostService - 포스트 작성 서비스
func WritePostService(body dto.WritePostDTO, ctx *gin.Context) (*app.ResponseException, error) {
	client := ctx.MustGet("client").(*ent.Client)

	tx, err := client.Tx(ctx)
	if err != nil {
		return app.TransactionsErrorResponse(err.Error(), nil), nil
	}

	store := NewPostStore(ctx, tx)

	userId, err := uuid.Parse(ctx.MustGet("id").(string))
	if err != nil {
		return app.UnAuthorizedErrorResponse("INVALID_USER_ID_UUID", nil), nil
	}

	user, err := tx.User.Query().Where(
		userEnt.IDEQ(userId),
	).First(ctx)

	if ent.IsNotFound(err) {
		return app.NotFoundErrorResponse("User Is Not Found", nil), nil
	}

	processedUrlSlug := body.UrlSlug
	urlSlugDuplicate, err := tx.Post.Query().Where(
		postEnt.And(
			postEnt.HasUserWith(
				userEnt.IDEQ(userId),
			),
			postEnt.URLSlug(body.UrlSlug),
		),
	).First(ctx)
	log.Println("post", urlSlugDuplicate)

	if !ent.IsNotFound(err) {
		processedUrlSlug = generateUrlSlug(body.Title)
	}

	if processedUrlSlug == "" {
		processedUrlSlug = generateUrlSlug(body.Title)
	}

	post, err := tx.Post.
		Create().
		SetTitle(body.Title).
		SetBody(body.Body).
		SetIsMarkdown(body.IsMarkdown).
		SetURLSlug(processedUrlSlug).
		SetIsTemp(body.IsTemp).
		SetIsPrivate(body.IsPrivate).
		SetThumbnail(body.Thumbnail).
		SetMeta(body.Meta).
		SetUserID(user.ID).
		SetUser(user).
		Save(ctx)

	// 포스트 생성이 실패한 경우
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return app.TransactionsErrorResponse(rerr.Error(), nil), nil
		}
		return app.InteralServerErrorResponse(err.Error(), nil), nil
	}

	log.Println("tags", body.Tags)
	tagObj, err := syncPostTags(post.ID, body.Tags, store)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return app.TransactionsErrorResponse(rerr.Error(), nil), nil
		}
		return app.InteralServerErrorResponse(err.Error(), nil), nil
	}

	// 태그 객체는 nil 일수도 있다. 그렇기 때문에 nil 체크를 한다.
	if tagObj != nil {
		missings := tagObj["missing"].([]uuid.UUID)
		addings := tagObj["adding"].([]uuid.UUID)

		log.Println("missings", missings)
		log.Println("adding", addings)

		if len(missings) > 0 {
			_, err := post.
				Update().
				RemoveTagIDs(missings...).
				Save(ctx)

			if err != nil {
				if rerr := tx.Rollback(); rerr != nil {
					return app.TransactionsErrorResponse(rerr.Error(), nil), nil
				}
				return app.InteralServerErrorResponse(err.Error(), nil), nil
			}
		}

		if len(addings) > 0 {
			_, err := post.
				Update().
				AddTagIDs(addings...).
				Save(ctx)

			if err != nil {
				if rerr := tx.Rollback(); rerr != nil {
					return app.TransactionsErrorResponse(rerr.Error(), nil), nil
				}
				return app.InteralServerErrorResponse(err.Error(), nil), nil
			}
		}
	}

	return &app.ResponseException{
		Code:          http.StatusOK,
		ResultCode:    0,
		Message:       "",
		ResultMessage: "",
		Data: libs.JSON{
			"id": post.ID,
		},
	}, tx.Commit()
}

func syncPostTags(postId uuid.UUID, tags []string, store *PostStore) (libs.JSON, error) {
	if len(tags) == 0 {
		return nil, nil
	}

	tx := store.tx
	bg := store.ctx

	currentTagList, err := tx.
		Post.
		Query().
		Where(postEnt.IDEQ(postId)).
		QueryTags().
		IDs(bg)

	log.Println("currentTags", currentTagList)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return nil, rerr
		}
		return nil, err
	}

	var tagList []uuid.UUID
	for _, tag := range tags {
		findTag, err := tx.Tag.
			Query().
			Where(tagEnt.NameEQ(tag)).
			First(bg)
		// 태그가 존재하지 않느 경우 Tag 를 생성
		if ent.IsNotFound(err) {
			createTag, err := tx.Tag.Create().
				SetName(tag).
				Save(bg)
			if err != nil {
				if rerr := tx.Rollback(); rerr != nil {
					return nil, rerr
				}
				return nil, err
			}
			tagList = append(tagList, createTag.ID)
		} else {
			// 존재하는 경우 해당 Tag ID 값을 가져온다
			tagList = append(tagList, findTag.ID)
		}
	}

	// 중복을 제거한 배열을 얻는다.
	var uniqueTagIds []uuid.UUID

	// uniqueTagIds 에 이미 존재하는 값인지 체크
	filterTagIds := make(map[uuid.UUID]bool)

	for _, value := range tagList {
		if _, tagId := filterTagIds[value]; !tagId {
			filterTagIds[value] = true
			uniqueTagIds = append(uniqueTagIds, value)
		}
	}

	log.Println("uniqueTagIds", uniqueTagIds)

	// 등록 할 태그
	var adding []uuid.UUID
	// 삭제 할 태그
	var missing []uuid.UUID

	if len(currentTagList) > len(uniqueTagIds) {
		for _, parent := range currentTagList {
			index := libs.FindUUID(uniqueTagIds, parent)

			if index < len(uniqueTagIds) && uniqueTagIds[index] == parent {
				// 이미 존재하는 경우
			} else {
				missing = append(missing, parent)
			}
		}

		for _, parent := range uniqueTagIds {
			index := libs.FindUUID(currentTagList, parent)

			if index < len(currentTagList) && currentTagList[index] == parent {
				// 이미 존재하는 경우
			} else {
				adding = append(adding, parent)
			}
		}
	} else {
		for _, parent := range uniqueTagIds {
			index := libs.FindUUID(currentTagList, parent)

			if index < len(currentTagList) && currentTagList[index] == parent {
				// 이미 존재하는 경우
			} else {
				adding = append(adding, parent)
			}
		}

		for _, parent := range currentTagList {
			index := libs.FindUUID(uniqueTagIds, parent)

			if index < len(uniqueTagIds) && uniqueTagIds[index] == parent {
				// 이미 존재하는 경우
			} else {
				missing = append(missing, parent)
			}
		}
	}

	return libs.JSON{
		"adding":  adding,
		"missing": missing,
	}, nil
}

func generateUrlSlug(title string) string {
	urlSlug := slug.Make(title)
	shortId := shortid.Generator()
	result := fmt.Sprintf("%v-%v", urlSlug, shortId.Generate())
	return result
}
