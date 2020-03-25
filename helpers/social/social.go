package social

import (
	"encoding/json"
	"golang.org/x/oauth2"
)

var redirectPath = "http://localhost:4000/api/v1.0/auth/social/callback/"

type SocialState struct {
	provider string
	next     string
}

type SocialAction interface {
	Github() string
	Facebook() string
	Google() string
}

func (s *SocialState) Google() string {
	callbackUri := redirectPath + "google"
	state, _ := json.Marshal(s.next)
	oauthConfing := &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  callbackUri,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile",
		},
	}
	return oauthConfing.AuthCodeURL(string(state))
}

func (s *SocialState) Facebook() string {
	facebookId := "ID"
	callbackUri := redirectPath + "facebook"
	state, _ := json.Marshal(s.next)
	return "https://www.facebook.com/v4.0/dialog/oauth?client_id=" + facebookId + "&redirect_uri=" + callbackUri + "&state=" + string(state) + "&scope=email,public_profile"
}

func (s *SocialState) Github() string {
	githubId := "ID"
	redirectUriWithNext := redirectPath + "github?next=" + s.next
	return "https://github.com/login/oauth/authorize?scope=user:email&client_id=" + githubId + "&redirect_uri=" + redirectUriWithNext
}

func Social(provider, next string) SocialAction {
	state := SocialState{
		provider: provider,
		next:     next,
	}
	return &state
}

func GenerateSocialLink(provier, next string) string {
	g := Social(provier, next)
	switch provier {
	case "facebook":
		return g.Facebook()
	case "google":
		return g.Google()
	case "github":
		return g.Github()
	default:
		return ""
	}
}
