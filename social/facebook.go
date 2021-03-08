package social

import (
	"encoding/json"
	"errors"
	"github.com/google/go-querystring/query"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

type Picture struct {
	Data struct {
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

type FacebookToken struct {
	AccessToken  string      `json:"access_token"`
	TokenType    string      `json:"token_type"`
	RefreshToken string      `json:"refresh_token"`
	Expiry       time.Time   `json:"expiry"`
	Raw          interface{} `json:"raw"`
}

type FacebookTokenJSON struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int32  `json:"expires_in"`
}

func (e *FacebookTokenJSON) expiry() (t time.Time) {
	if v := e.ExpiresIn; v != 0 {
		return time.Now().Add(time.Duration(v) * time.Second)
	}
	return
}

func GetFacebookAccessToken(code string) string {
	id := os.Getenv("FACEBOOK_CLIENT_ID")
	secret := os.Getenv("FACEBOOK_CLIENT_SECRET")
	client := &http.Client{Timeout: 10 * time.Second}

	type FacebookOAuthParams struct {
		Code         string `url:"code" json:"code"`
		ClientID     string `url:"client_id" json:"client_id"`
		ClientSecret string `url:"client_secret" json:"client_secret"`
		RedirectUri  string `url:"redirect_uri" json:"redirect_uri"`
	}

	params := FacebookOAuthParams{
		Code:         code,
		ClientID:     id,
		ClientSecret: secret,
		RedirectUri:  "http://localhost:8080/api/v1.0/auth/social/callback/facebook",
	}

	queryString, _ := query.Values(params)
	req, err := http.NewRequest("GET", "https://graph.facebook.com/v4.0/oauth/access_token?"+queryString.Encode(), nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1<<20))
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	var token *FacebookToken
	content, _, _ := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	switch content {
	case "application/x-www-form-urlencoded", "text/plain":
		values, err := url.ParseQuery(string(body))
		if err != nil {
			panic(err)
		}

		token = &FacebookToken{
			AccessToken:  values.Get("access_token"),
			TokenType:    values.Get("token_type"),
			RefreshToken: values.Get("refresh_token"),
			Raw:          values,
		}
		e := values.Get("expires_in")
		expires, _ := strconv.Atoi(e)
		if expires != 0 {
			token.Expiry = time.Now().Add(time.Duration(expires) * time.Second)
		}
	default:
		var tokenJson FacebookTokenJSON
		if err = json.Unmarshal(body, &tokenJson); err != nil {
			panic(err)
		}

		token = &FacebookToken{
			AccessToken:  tokenJson.AccessToken,
			TokenType:    tokenJson.TokenType,
			RefreshToken: tokenJson.RefreshToken,
			Expiry:       tokenJson.expiry(),
			Raw:          make(map[string]interface{}),
		}
		json.Unmarshal(body, &token.Raw)
	}

	if token.AccessToken == "" {
		panic(errors.New("oauth2: server response missing access_token"))
	}
	return token.AccessToken
}

func GetFacebookProfile(token string) *SocialProfile {
	req, err := http.NewRequest("GET", "https://graph.facebook.com/v4.0/me?fields=id,name,email,picture", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
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

	profile := SocialProfile{
		ID:        result.ID,
		Name:      result.Name,
		Email:     result.Email,
		Thumbnail: result.Picture.Data.Url,
	}

	return &profile
}
