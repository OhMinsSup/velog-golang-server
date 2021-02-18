package social

import (
	"encoding/json"
	"fmt"
	"github.com/OhMinsSup/story-server/libs"
	"golang.org/x/oauth2"
)

var redirectPath = "http://localhost:8080/api/v1.0/auth/social/callback/"

type SocialState struct {
	Provider string `json:"provider"`
	Next     string `json:"next"`
}

type SocialAction interface {
	Github() string
	Facebook() string
	Google() string
	Kakao() string
}

func (s SocialState) Google() string {
	callbackUri := redirectPath + "google"
	state, _ := json.Marshal(s.Next)
	oauthConfig := &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  callbackUri,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile",
		},
	}
	return oauthConfig.AuthCodeURL(string(state))
}

func (s SocialState) Facebook() string {
	id := libs.GetEnvWithKey("FACEBOOK_CLIENT_ID")
	redirectUrl := redirectPath + "facebook"
	state, _ := json.Marshal(s.Next)
	return fmt.Sprintf("https://www.facebook.com/v4.0/dialog/oauth?client_id=%v&redirect_uri=%v&state=%v&scope=email,public_profile", id, redirectUrl, state)
}

func (s SocialState) Github() string {
	id := libs.GetEnvWithKey("GITHUB_CLIENT_ID")
	redirectUriWithNext := redirectPath + "github?next=" + s.Next
	return fmt.Sprintf("https://github.com/login/oauth/authorize?scope=user:email&client_id=%v&redirect_uri=%v", id, redirectUriWithNext)
}

func (s SocialState) Kakao() string {
	restKey := libs.GetEnvWithKey("KAKAO_REST_API_KEY")
	state, _ := json.Marshal(s.Next)
	redirectUrl := redirectPath + "kakao"
	return fmt.Sprintf("https://kauth.kakao.com/oauth/authorize?response_type=code&client_id=%v&redirect_uri=%v&state=%v", restKey, redirectUrl, state)
}

func Social(provider, next string) SocialAction {
	state := SocialState{
		Provider: provider,
		Next:     next,
	}
	return &state
}

func GenerateSocialLink(provider, next string) string {
	snapshot := Social(provider, next)
	switch provider {
	case "facebook":
		return snapshot.Facebook()
	case "google":
		return snapshot.Google()
	case "github":
		return snapshot.Github()
	case "kakao":
		return snapshot.Kakao()
	default:
		return ""
	}
}
