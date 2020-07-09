package repository

import (
	"github.com/OhMinsSup/story-server/models"
	"github.com/SKAhack/go-shortid"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (a *AuthRepository) ExistEmail(email string) (bool, int, error) {
	var user models.User
	err := a.db.Where("email = ?", strings.ToLower(email)).First(&user).Error
	if !gorm.IsRecordNotFoundError(err) {
		return true, http.StatusOK, nil
	} else {
		return false, http.StatusOK, nil
	}
}

func (a *AuthRepository) CreateEmailAuth(email string) (*models.EmailAuth, int, error) {
	shortId := shortid.Generator()

	tx := a.db.Begin()
	emailAuth := models.EmailAuth{
		Email: email,
		Code:  shortId.Generate(),
	}

	if err := tx.Create(&emailAuth).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, err
	}

	return &emailAuth, http.StatusOK, nil
}
