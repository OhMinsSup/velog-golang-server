package social

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"net/http"
	"time"
)

var (
	facebook_clientId     = ""
	facebook_clientSecret = ""
)

type Picture struct {
	Data map[string]struct {
		Height       int    `json:"height"`
		IsSilhouette bool   `json:"is_silhouette"`
		Url          string `json:"url"`
		Width        int    `json:"width"`
	} `json:"data"`
}

type FacebookProfile struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Picture Picture `json:"picture"`
}

type FacebookOAuthResult struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
}

type FacebookOAuthParams struct {
	Code         string `url:"code"`
	ClientID     string `url:"client_id"`
	ClientSecret string `url:"client_secret"`
	RedirectUri  string `url: "redirect_uri"`
}

func GetFacebookAccessToken(code, redirectUri string) string {
	oauthParams := FacebookOAuthParams{Code: code, ClientID: facebook_clientId, ClientSecret: facebook_clientSecret, RedirectUri: redirectUri}
	queryString, _ := query.Values(oauthParams)
	req, err := http.NewRequest("GET", "https://graph.facebook.com/v4.0/oauth/access_token?"+queryString.Encode(), nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var result FacebookOAuthResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		panic(err)
	}
	return result.AccessToken
}

func GetFacebookProfile(token string) FacebookProfile {
	req, err := http.NewRequest("GET", "https://graph.facebook.com/v4.0/me?fields=id,name,email,picture", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+token)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var result FacebookProfile
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		panic(err)
	}
	return result
}
