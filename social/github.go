package social

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type GithubTokenSource struct {
	AccessToken string `json:"access_token"`
}

type GithubToken struct {
	AccessToken  string
	TokenType    string
	RefreshToken string
	Expiry       time.Time
	Raw          interface{}
}

type GithubTokenJSON struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int32  `json:"expires_in"`
}

func GetGithubAccessToken(code string) string {
	id := os.Getenv("GITHUB_CLIENT_ID")
	secret := os.Getenv("GITHUB_CLIENT_SECRET")
	client := http.Client{Timeout: 10 * time.Second}

	v := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"client_secret": {secret},
		"client_id":     {id},
	}

	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", strings.NewReader(v.Encode()))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1<<20))
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	var token *GithubToken
	content, _, _ := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	switch content {
	case "application/x-www-form-urlencoded", "text/plain":
		values, err := url.ParseQuery(string(body))
		if err != nil {
			panic(err)
		}

		token = &GithubToken{
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
		var tokenJson GithubTokenJSON
		if err = json.Unmarshal(body, &tokenJson); err != nil {
			panic(err)
		}

		token = &GithubToken{
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

func (e *GithubTokenJSON) expiry() (t time.Time) {
	if v := e.ExpiresIn; v != 0 {
		return time.Now().Add(time.Duration(v) * time.Second)
	}
	return
}

func (t *GithubTokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func GetGithubProfile(accessToken string) *SocialProfile {
	tokenSource := &GithubTokenSource{
		AccessToken: accessToken,
	}

	ctx := context.Background()
	oauthClient := oauth2.NewClient(ctx, tokenSource)
	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get(ctx, "")
	if err != nil {
		panic(err)
	}

	profile := SocialProfile{
		ID:        strconv.FormatInt(*user.ID, 10),
		Name:      *user.Name,
		Email:     *user.Email,
		Thumbnail: *user.AvatarURL,
	}

	return &profile
}
