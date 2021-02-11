package social

import (
	"encoding/json"
	"golang.org/x/oauth2"
	"os"
)

var redirectPath = "http://localhost:8080/api/v1.0/auth/social/callback/"

type State struct {
	Provider string `json:"provider"`
	Next     string `json:"next"`
}

type Action interface {
	Github() string
	Facebook() string
	Google() string
}

func (s State) Google() string {
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

func (s State) Facebook() string {
	id := os.Getenv("FACEBOOK_CLIENT_ID")
	callbackUri := redirectPath + "facebook"
	state, _ := json.Marshal(s.Next)
	return "https://www.facebook.com/v4.0/dialog/oauth?client_id=" + id + "&redirect_uri=" + callbackUri + "&state=" + string(state) + "&scope=email,public_profile"
}

func (s State) Github() string {
	id := os.Getenv("GITHUB_CLIENT_ID")
	redirectUriWithNext := redirectPath + "github?next=" + s.Next
	return "https://github.com/login/oauth/authorize?scope=user:email&client_id=" + id + "&redirect_uri=" + redirectUriWithNext
}

func (s State) Kakao() string  {
	return ""
}

func Social(provider, next string) Action {
	state := State{
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
	default:
		return ""
	}
}
