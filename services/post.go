package services

import (
	"context"
	"fmt"
	"github.com/OhMinsSup/story-server/app"
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/ent"
	postEnt "github.com/OhMinsSup/story-server/ent/post"
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
			postEnt.FkUserID(userId),
			postEnt.URLSlug(body.UrlSlug),
		),
	).First(bg)
	log.Println("post", urlSlugDuplicate)

	if !ent.IsNotFound(err) {
		if processedUrlSlug == "" {
			processedUrlSlug = slug.Make(body.Title)
		}

		processedUrlSlug = fmt.Sprintf("%v-%v", processedUrlSlug, shortid.Generator())
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
		SetFkUserID(user.ID).
		SetUser(user).
		Save(bg)

	// 포스트 생성이 실패한 경
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
