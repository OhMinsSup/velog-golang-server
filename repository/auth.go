package repository

import (
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/libs"
	"github.com/OhMinsSup/story-server/models"
	"github.com/SKAhack/go-shortid"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

// AuthRepository
type AuthRepository struct {
	db *gorm.DB
}

// NewAuthRepository auth 저장소를 생성해서 auth 와 관련된 기능들을 제공
func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

// ExistEmail 이메일이 현재 db에 등록되어 있는지 체크
func (a *AuthRepository) ExistEmail(email string) (bool, int, error) {
	// 이메일이 존재하지 않는 경우
	if email == "" {
		return false, http.StatusBadRequest, errors.New("Email is Empty")
	}

	var user models.User
	err := a.db.Where("email = ?", strings.ToLower(email)).First(&user).Error

	if !gorm.IsRecordNotFoundError(err) {
		// 이메일이 존쟈하는 경우 (로그인)
		return true, http.StatusOK, nil
	} else {
		// 이메일이 존재하지 않는 경우 (회원가입)
		return false, http.StatusOK, nil
	}
}

// ExistsCode 이메일 인증 코드 유효성 체크
func (a *AuthRepository) ExistsCode(code string) (*models.EmailAuth, int, error) {
	if code == "" {
		return nil, http.StatusBadRequest, errors.New("Code is Empty")
	}

	var emailAuth models.EmailAuth
	err := a.db.Where("code = ?", code).First(&emailAuth).Error
	// 코드가 존재하지 않는경우 badRequest
	if gorm.IsRecordNotFoundError(err) {
		// 코드가 존재하지 않는 경우는 db 테이블에 데이터가 존재하지 않는 경우
		return nil, http.StatusBadRequest, err
	} else {
		// 존재하는 경우에는 emailAuth 정보를 리턴
		return &emailAuth, http.StatusOK, nil
	}
}

// CreateEmailAuth 이메일 인증 모델 생성
func (a *AuthRepository) CreateEmailAuth(email string) (*models.EmailAuth, int, error) {
	// 이메일이 존재하지 않는 경우
	if email == "" {
		return nil, http.StatusBadRequest, errors.New("Email is Empty")
	}

	// 인증 code Id 값
	shortId := shortid.Generator()

	tx := a.db.Begin()

	// 이메일 인증 모델 생성
	emailAuth := models.EmailAuth{
		Email: email,
		Code:  shortId.Generate(),
	}

	// 이메일 인증 모델 실제 db에 생성
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
		return nil, http.StatusConflict, libs.ErrorAlreadyExists
	}

	return &user, http.StatusOK, nil
}

// ExistsByEmailAndUsername - 이메일및 유저명이 이미 존재하는지 존재하지않는지 체크
func (a *AuthRepository) ExistsByEmailAndUsername(username, email string) (bool, int, error) {
	var user models.User
	err := a.db.Where("email = ?", email).Or("username = ?", username).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return true, http.StatusOK, nil
	}

	return false, http.StatusConflict, libs.ErrorAlreadyExists
}

func (a *AuthRepository) SocialUser(userData dto.SocialUserParams) (*models.User, *models.UserProfile, int, error) {
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

	return &user, &userProfile, http.StatusOK, tx.Commit().Error
}

// CreateUser - 유저가 회원가입을 할 경우 유저 생성
func (a *AuthRepository) CreateUser(userData dto.LocalRegisterDTO) (*models.User, int, error) {
	tx := a.db.Begin()
	user := models.User{
		Email:       userData.Email,
		IsCertified: true,
		Username:    userData.Username,
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, err
	}

	userProfile := models.UserProfile{
		DisplayName: userData.DisplayName,
		ShortBio:    userData.ShortBio,
		UserID:      user.ID,
	}

	if err := tx.Create(&userProfile).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, err
	}

	velogConfig := models.VelogConfig{
		UserID: user.ID,
	}

	if err := tx.Create(&velogConfig).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, err
	}

	userMeta := models.UserMeta{
		UserID: user.ID,
	}

	if err := tx.Create(&userMeta).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, err
	}

	return &user, http.StatusOK, tx.Commit().Error
}
