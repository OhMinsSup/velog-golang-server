package social

import (
	"encoding/json"
	"errors"
	"github.com/OhMinsSup/story-server/libs"
	"github.com/google/go-querystring/query"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type KakaoProfile struct {
	ID           int64        `json:"id"`
	KakaoAccount KakaoAccount `json:"kakao_account"`
}

type KakaoAccount struct {
	ProfileNeedsAgreement  bool               `json:"profile_needs_agreement"`
	Profile                KakaoDetailProfile `json:"profile"`
	EmailNeedsAgreement    bool               `json:"email_needs_agreement"`
	IsEmailValid           bool               `json:"is_email_valid"`
	IsEmailVerified        bool               `json:"is_email_verified"`
	Email                  string             `json:"email"`
	AgeRangeNeedsAgreement bool               `json:"age_range_needs_agreement"`
	AgeRange               string             `json:"age_range"`
	BirthdayNeedsAgreement bool               `json:"birthday_needs_agreement"`
	Birthday               string             `json:"birthday"`
	GenderNeedsAgreement   bool               `json:"gender_needs_agreement"`
	Gender                 string             `json:"gender"`
}

type KakaoDetailProfile struct {
	Nickname          string `json:"nickname"`
	ThumbnailImageURL string `json:"thumbnail_image_url"`
	ProfileImageURL   string `json:"profile_image_url"`
}

type KakaoToken struct {
	TokenType             string    `json:"token_type"`
	AccessToken           string    `json:"access_token"`
	RefreshToken          string    `json:"refresh_token"`
	ExpiresIn             time.Time `json:"expires_in"`
	Scope                 string    `json:"scope"`
	RefreshTokenExpiresIn time.Time `json:"refresh_token_expires_in"`
}

type KakaoTokenJSON struct {
	TokenType             string `json:"token_type"`
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	ExpiresIn             int64  `json:"expires_in"`
	Scope                 string `json:"scope"`
	RefreshTokenExpiresIn int64  `json:"refresh_token_expires_in"`
}

func (e *KakaoTokenJSON) expiry() (t time.Time) {
	if v := e.ExpiresIn; v != 0 {
		return time.Now().Add(time.Duration(v) * time.Second)
	}
	return
}

func (e *KakaoTokenJSON) refreshExpiry() (t time.Time) {
	if v := e.RefreshTokenExpiresIn; v != 0 {
		return time.Now().Add(time.Duration(v) * time.Second)
	}
	return
}

func GetKakaoAccessToken(code string) string {
	id := libs.GetEnvWithKey("KAKAO_REST_API_KEY")
	secret := libs.GetEnvWithKey("KAKAO_CLIENT_SECRET")
	client := &http.Client{Timeout: 10 * time.Second}

	type KakaoOAuthParams struct {
		Code         string `url:"code" json:"code"`
		ClientID     string `url:"client_id" json:"client_id"`
		ClientSecret string `url:"client_secret" json:"client_secret"`
		RedirectUri  string `url:"redirect_uri" json:"redirect_uri"`
		GrantType    string `url:"grant_type" json:"grant_type"`
	}

	params := KakaoOAuthParams{
		Code:         code,
		ClientID:     id,
		ClientSecret: secret,
		GrantType:    "authorization_code",
		RedirectUri:  "http://localhost:8080/api/v1.0/auth/social/callback/kakao",
	}

	queryString, _ := query.Values(params)
	req, err := http.NewRequest("POST", "https://kauth.kakao.com/oauth/authorize?"+queryString.Encode(), nil)
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

	var token *KakaoToken
	content, _, _ := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	switch content {
	case "application/x-www-form-urlencoded", "text/plain":
		values, err := url.ParseQuery(string(body))
		if err != nil {
			panic(err)
		}
		log.Println("logs", values)
		token = &KakaoToken{
			AccessToken:  values.Get("access_token"),
			TokenType:    values.Get("token_type"),
			RefreshToken: values.Get("refresh_token"),
			Scope:        values.Get("scope"),
		}

		e := values.Get("expires_in")
		expires, _ := strconv.Atoi(e)
		if expires != 0 {
			token.ExpiresIn = time.Now().Add(time.Duration(expires) * time.Second)
		}

		re := values.Get("refresh_token_expires_in")
		refreshExpires, _ := strconv.Atoi(re)
		if refreshExpires != 0 {
			token.RefreshTokenExpiresIn = time.Now().Add(time.Duration(refreshExpires) * time.Second)
		}
	default:
		var tokenJson KakaoTokenJSON
		if err = json.Unmarshal(body, &tokenJson); err != nil {
			panic(err)
		}
		token = &KakaoToken{
			AccessToken:           tokenJson.AccessToken,
			TokenType:             tokenJson.TokenType,
			RefreshToken:          tokenJson.RefreshToken,
			Scope:                 tokenJson.Scope,
			ExpiresIn:             tokenJson.expiry(),
			RefreshTokenExpiresIn: tokenJson.refreshExpiry(),
		}
	}

	if token.AccessToken == "" {
		panic(errors.New("oauth2: server response missing access_token"))
	}
	return token.AccessToken
}

func GetKakaoProfile(token string) *SocialProfile {
	req, err := http.NewRequest("POST", "https://kapi.kakao.com/v2/user/me", nil)
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

	var result KakaoProfile
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		panic(err)
	}

	profile := SocialProfile{
		ID:        strconv.FormatInt(result.ID, 10),
		Name:      result.KakaoAccount.Profile.Nickname,
		Email:     result.KakaoAccount.Email,
		Thumbnail: result.KakaoAccount.Profile.ThumbnailImageURL,
	}

	return &profile
}
