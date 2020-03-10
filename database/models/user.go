package models

import "time"

type User struct {
	ID          string `gorm:"primary_key"`
	Username    string `sql:"index"`
	Email       string `sql:"index"`
	IsCertified bool   `gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}


func (user *User) GenerateUserToken() {

}

func (user *User) RefreshUserToken() {

}
