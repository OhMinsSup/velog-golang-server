package social

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
	"time"
)

type GithubOAuthResult struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type GithubOAuthParams struct {
	Code         string `json:"code"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func GetGithubAccessToken(code string) string {
	id := os.Getenv("GITHUB_CLIENT_ID")
	secret := os.Getenv("GITHUB_CLIENT_SECRET")

	oauthParams := GithubOAuthParams{code, id, secret}
	data, _ := json.Marshal(oauthParams)
	buff := bytes.NewBuffer(data)
	client := http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", buff)

	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var result GithubOAuthResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		panic(err)
	}

	return result.AccessToken
}

type TokenSource struct {
	AccessToken string `json:"access_token"`
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func GetGithubProfile(accessToken string) *github.User {
	tokenSource := &TokenSource{
		AccessToken: accessToken,
	}

	ctx := context.Background()
	oauthClient := oauth2.NewClient(ctx, tokenSource)
	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get(ctx, "")
	log.Println("user:", user)
	if err != nil {
		panic(err)
	}
	return user
}
