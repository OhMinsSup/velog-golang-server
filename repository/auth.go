package repository

import (
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/models"
	"github.com/SKAhack/go-shortid"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
)

type AuthRepository struct {
	db *gorm.DB
}

type CreateUserParams struct {
	Email       string `json:"email"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	ShortBio    string `json:"short_bio"`
	UserID      string `json:"user_id"`
}

type SocialUserParams struct {
	Email       string `json:"email"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	ShortBio    string `json:"short_bio"`
	UserID      string `json:"user_id"`
	AccessToken string `json:"access_token"`
	Provider    string `json:"provider"`
	SocialID    string `json:"social_id"`
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

// 이메일이 현재 db에 등록되어 있는지 체크
func (a *AuthRepository) ExistEmail(email string) (bool, int, error) {
	var user models.User
	err := a.db.Where("email = ?", strings.ToLower(email)).First(&user).Error

	if !gorm.IsRecordNotFoundError(err) {
		// 이메일이 존재하지 않는 경우 회원가입
		return true, http.StatusOK, nil
	} else {
		// 이메일이 존쟈하는 경우 로그인
		return false, http.StatusOK, nil
	}
}

// 이메일 인증 코드 유효성 체크
func (a *AuthRepository) ExistsCode(code string) (*models.EmailAuth, int, error) {
	var emailAuth models.EmailAuth
	err := a.db.Where("code = ?", code).First(&emailAuth).Error
	// 코드가 존재하지 않는경우 badRequest
	if gorm.IsRecordNotFoundError(err) {
		return nil, http.StatusBadRequest, err
	} else {
		return &emailAuth, http.StatusOK, nil
	}
}

// 이메일 인증 모델 생성
func (a *AuthRepository) CreateEmailAuth(email string) (*models.EmailAuth, int, error) {
	shortId := shortid.Generator()

	tx := a.db.Begin()

	// 이메일 인증 모델 생성
	emailAuth := models.EmailAuth{
		Email: email,
		Code:  shortId.Generate(),
	}

	// 이메일 생성
	if err := tx.Create(&emailAuth).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, err
	}

	return &emailAuth, http.StatusOK, tx.Commit().Error
}

func (a *AuthRepository) FindByEmailAndUsername(username, email string) (*models.User, int, error) {
	var user models.User
	err := a.db.Where("email = ?", email).Or("username = ?", username).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, http.StatusConflict, helpers.ErrorAlreadyExists
	}

	return &user, http.StatusOK, nil
}

func (a *AuthRepository) ExistsByEmailAndUsername(username, email string) (bool, int, error) {
	var user models.User
	err := a.db.Where("email = ?", email).Or("username = ?", username).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return true, http.StatusOK, nil
	}

	return false, http.StatusConflict, helpers.ErrorAlreadyExists
}

func (a *AuthRepository) SocialUser(userData SocialUserParams) (*models.User, *models.UserProfile, int, error) {
	tx := a.db.Begin()
	user := models.User{
		Email:       userData.Email,
		IsCertified: true,
		Username:    userData.Username,
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, nil, http.StatusInternalServerError, err
	}

	socialAccount := models.SocialAccount{
		AccessToken: userData.AccessToken,
		Provider:    userData.Provider,
		UserID:      user.ID,
		SocialId:    userData.SocialID,
	}

	if err := tx.Create(&socialAccount).Error; err != nil {
		tx.Rollback()
		return nil, nil, http.StatusInternalServerError, err
	}

	userProfile := models.UserProfile{
		DisplayName: userData.DisplayName,
		ShortBio:    userData.ShortBio,
		UserID:      user.ID,
	}

	if err := tx.Create(&userProfile).Error; err != nil {
		tx.Rollback()
		return nil, nil, http.StatusInternalServerError, err
	}

	velogConfig := models.VelogConfig{
		UserID: user.ID,
	}

	if err := tx.Create(&velogConfig).Error; err != nil {
		tx.Rollback()
		return nil, nil, http.StatusInternalServerError, err
	}

	userMeta := models.UserMeta{
		UserID: user.ID,
	}

	if err := tx.Create(&userMeta).Error; err != nil {
		tx.Rollback()
		return nil, nil, http.StatusInternalServerError, err
	}

	return  &user, &userProfile, http.StatusOK, tx.Commit().Error
}

func (a *AuthRepository) CreateUser(userData CreateUserParams) (*models.User, *models.UserProfile, int, error) {
	tx := a.db.Begin()
	user := models.User{
		Email:       userData.Email,
		IsCertified: true,
		Username:    userData.Username,
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, nil, http.StatusInternalServerError, err
	}

	userProfile := models.UserProfile{
		DisplayName: userData.DisplayName,
		ShortBio:    userData.ShortBio,
		UserID:      user.ID,
	}

	if err := tx.Create(&userProfile).Error; err != nil {
		tx.Rollback()
		return nil, nil, http.StatusInternalServerError, err
	}

	velogConfig := models.VelogConfig{
		UserID: user.ID,
	}

	if err := tx.Create(&velogConfig).Error; err != nil {
		tx.Rollback()
		return nil, nil, http.StatusInternalServerError, err
	}

	userMeta := models.UserMeta{
		UserID: user.ID,
	}

	if err := tx.Create(&userMeta).Error; err != nil {
		tx.Rollback()
		return nil, nil, http.StatusInternalServerError, err
	}

	return &user, &userProfile, http.StatusOK, tx.Commit().Error
}
