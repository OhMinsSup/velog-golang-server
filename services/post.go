package services

import (
	"context"
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
)

// 공용 스토어
type PostStore struct {
	tx  *ent.Tx
	ctx context.Context
}

func NewPostStore(ctx context.Context, tx *ent.Tx) *PostStore {
	return &PostStore{
		tx:  tx,
		ctx: ctx,
	}
}

// UpdatePostService - 포스트 수정 서비스
func UpdatePostService(body dto.UpdatePostDTO, ctx *gin.Context) (*app.ResponseException, error) {
	client := ctx.MustGet("client").(*ent.Client)
	bg := context.Background()

	tx, err := client.Tx(bg)
	if err != nil {
		return app.TransactionsErrorResponse(err.Error(), nil), nil
	}

	store := NewPostStore(bg, tx)
	log.Println(store)

	userId, err := uuid.Parse(ctx.MustGet("id").(string))
	if err != nil {
		return app.UnAuthorizedErrorResponse("INVALID_USER_ID_UUID", nil), nil
	}

	user, err := tx.User.Query().Where(
		userEnt.IDEQ(userId),
	).First(bg)

	log.Println(user)
	if ent.IsNotFound(err) {
		return app.NotFoundErrorResponse("User Is Not Found", nil), nil
	}

	return &app.ResponseException{
		Code:          http.StatusOK,
		ResultCode:    0,
		Message:       "",
		ResultMessage: "",
		Data: libs.JSON{
			"id": true,
		},
	}, nil
}

// WritePostService - 포스트 작성 서비스
func WritePostService(body dto.WritePostDTO, ctx *gin.Context) (*app.ResponseException, error) {
	client := ctx.MustGet("client").(*ent.Client)
	bg := context.Background()

	tx, err := client.Tx(bg)
	if err != nil {
		return app.TransactionsErrorResponse(err.Error(), nil), nil
	}

	store := NewPostStore(bg, tx)

	userId, err := uuid.Parse(ctx.MustGet("id").(string))
	if err != nil {
		return app.UnAuthorizedErrorResponse("INVALID_USER_ID_UUID", nil), nil
	}

	user, err := tx.User.Query().Where(
		userEnt.IDEQ(userId),
	).First(bg)

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
	).First(bg)
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
		Save(bg)

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
				Save(bg)

			if err != nil {
				log.Println("err missing", err)
				if rerr := tx.Rollback(); rerr != nil {
					log.Println("rerr missing", rerr)
					return app.TransactionsErrorResponse(rerr.Error(), nil), nil
				}
				return app.InteralServerErrorResponse(err.Error(), nil), nil
			}
		}

		if len(addings) > 0 {
			_, err := post.
				Update().
				AddTagIDs(addings...).
				Save(bg)

			if err != nil {
				log.Println("err missing", err)
				if rerr := tx.Rollback(); rerr != nil {
					log.Println("rerr missing", rerr)
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

	log.Println("errors", err)
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

	log.Println("tagList", tagList)
	// 중복을 제거한 배열을 얻는다.
	var uniqueTagIds []uuid.UUID

	// uniqueTagIds에 이미 존재하는 값인지 체크
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

	// 이미 존재하는 값인지 체크
	comparesMissing := make(map[uuid.UUID]bool)
	comparesAdding := make(map[uuid.UUID]bool)

	if len(currentTagList) > len(uniqueTagIds) {
		// ["1", "2"] 추가 할 태그
		for _, parent := range uniqueTagIds {
			// ["3", "4", "5", "1"] 이미 등록된 태그
			for _, children := range currentTagList {
				// 이미 등록된 태그에 똑같은 태그값을 넣은 경우에는 그냥 넘긴다.
				if parent == children {
					break
				}

				// 추가 할 태그와 이미 등록된 태그가 같지 않으면 중복되지 않은 이미 등록된 값들은
				// missing 에 넣어준다
				if parent != children && !comparesMissing[children] {
					comparesMissing[children] = true
					missing = append(missing, children)
				}
			}
		}

		// ["3", "4", "5", "1"] 이미 등록된 태그
		for _, parent := range currentTagList {
			// ["1", "2"] 추가 할 태그
			for _, children := range uniqueTagIds {
				// 이미 등록된 태그에 똑같은 태그값을 넣은 경우에는 그냥 넘긴다.
				if parent == children {
					break
				}

				// 추가 할 태그와 이미 등록된 태그가 같지 않으면 중복되지 않은 이미 등록된 값들은
				// adding 에 넣어준다
				if parent != children && !comparesAdding[children] {
					comparesAdding[children] = true
					adding = append(adding, children)
				}
			}
		}
	} else {
		// ["1", "2"]  이미 등록된 태그
		for _, parent := range currentTagList {
			// ["3", "4", "5", "1"] 추가 할 태그
			for _, children := range uniqueTagIds {
				// 이미 등록된 태그에 똑같은 태그값을 넣은 경우에는 그냥 넘긴다.
				if parent == children {
					break
				}

				// 추가 할 태그와 이미 등록된 태그가 같지 않으면 중복되지 않은 이미 등록된 값들은
				// missing 에 넣어준다
				if parent != children && !comparesMissing[children] {
					comparesMissing[children] = true
					missing = append(missing, children)
				}
			}
		}

		// ["3", "4", "5", "1"] 추가 할 태그
		for _, parent := range uniqueTagIds {
			if len(currentTagList) == 0 {
				// 이미 등록한 태그가 존재하지 않는 경우
				adding = append(adding, parent)
			} else {
				// 이미 등록한 태그가 존재하는 경
				// ["1", "2"]  이미 등록된 태그
				for _, children := range currentTagList {
					// 이미 등록된 태그에 똑같은 태그값을 넣은 경우에는 그냥 넘긴다.
					if parent == children {
						break
					}

					// 추가 할 태그와 이미 등록된 태그가 같지 않으면 중복되지 않은 이미 등록된 값들은
					// adding 에 넣어준다
					if parent != children && !comparesAdding[children] {
						comparesAdding[children] = true
						adding = append(adding, children)
					}
				}
			}
		}
	}

	log.Println("adding", adding)
	log.Println("missing", missing)

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
