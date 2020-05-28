package social

import (
	"encoding/json"
	"golang.org/x/oauth2"
)

var redirectPath = "http://localhost:4000/api/v1.0/auth/social/callback/"

type State struct {
	provider string `json:"provider"`
	next     string `json:"next"`
}

type Action interface {
	Github() string
	Facebook() string
	Google() string
}

func (s *State) Google() string {
	callbackUri := redirectPath + "google"
	state, _ := json.Marshal(s.next)
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

func (s *State) Facebook() string {
	facebookId := "ID"
	callbackUri := redirectPath + "facebook"
	state, _ := json.Marshal(s.next)
	return "https://www.facebook.com/v4.0/dialog/oauth?client_id=" + facebookId + "&redirect_uri=" + callbackUri + "&state=" + string(state) + "&scope=email,public_profile"
}

func (s *State) Github() string {
	githubId := "ID"
	redirectUriWithNext := redirectPath + "github?next=" + s.next
	return "https://github.com/login/oauth/authorize?scope=user:email&client_id=" + githubId + "&redirect_uri=" + redirectUriWithNext
}

func Social(provider, next string) Action {
	state := State{
		provider: provider,
		next:     next,
	}
	return &state
}

func GenerateSocialLink(provider, next string) string {
	uri := Social(provider, next)
	switch provider {
	case "facebook":
		return uri.Facebook()
	case "google":
		return uri.Google()
	case "github":
		return uri.Github()
	default:
		return ""
	}
}
