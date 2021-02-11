package services

import (
	"context"
	"github.com/OhMinsSup/story-server/app"
	"github.com/OhMinsSup/story-server/ent"
	socialaccountEnt "github.com/OhMinsSup/story-server/ent/socialaccount"
	"github.com/OhMinsSup/story-server/helpers/social"
	"github.com/gin-gonic/gin"
	"log"
)

func getSocialData(client *ent.Client, provider, code string) {
	var accessToken string
	switch provider {
	case "facebook":
		accessToken = social.GetFacebookAccessToken(code)
		break
	case "github":
		accessToken = social.GetGithubAccessToken(code)
		break
	case "kakao":
		accessToken = ""
		break
	case "google":
		accessToken = ""
		break
	}

	bg := context.Background()

	socialAccount, err := client.SocialAccount.
		Query().
		Where(
			socialaccountEnt.And(
				socialaccountEnt.SocialIDEQ(""),
				socialaccountEnt.ProviderEQ(provider))).First(bg)

	log.Println(accessToken)
}

func SocialCallbackService(ctx *gin.Context) (*app.ResponseException, error) {
	code := ctx.Query("code")
	if code == "" {
		return app.BadRequestErrorResponse("CODE IS EMPTY", nil), nil
	}

	return nil, nil
}
