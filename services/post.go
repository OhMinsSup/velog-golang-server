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

func WritePostService(body dto.WritePostDTO, ctx *gin.Context) (*app.ResponseException, error) {
	client := ctx.MustGet("client").(*ent.Client)
	bg := context.Background()

	tx, err := client.Tx(bg)
	if err != nil {
		return app.TransactionsErrorResponse(err.Error(), nil), nil
	}

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

func syncPostTags(postId uuid.UUID, tags []string, tx *ent.Tx) (*[]ent.Tag, error) {
	bg := context.Background()

	var tagIds []uuid.UUID
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
			tagIds = append(tagIds, createTag.ID)
		} else {
			// 존재하는 경우 해당 Tag ID 값을 가져온다
			tagIds = append(tagIds, findTag.ID)
		}
	}



	return nil, nil
}

func generateUrlSlug(title string) string {
	urlSlug := slug.Make(title)
	shortId := shortid.Generator()
	result := fmt.Sprintf("%v-%v", urlSlug, shortId.Generate())
	return result
}
