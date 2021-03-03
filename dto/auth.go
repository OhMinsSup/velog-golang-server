package dto

import "github.com/OhMinsSup/story-server/libs/social"

// SendEmailBody - SendEmailController 이메일 발송 body 데이터
type SendEmailBody struct {
	Email string `json:"email" binding:"email,required"`
}

// CodeParams - CodeController 코드 인증에 관한 params 데이터
type CodeParams struct {
	Code string `json:"code"`
}

// LocalRegisterBody - LocalRegisterController 회원가입 body 데이터
type LocalRegisterBody struct {
	RegisterToken string `json:"register_token" binding:"required"`
	DisplayName   string `json:"display_name" binding:"required"`
	UserName      string `json:"username" binding:"required"`
	ShortBio      string `json:"short_bio"`
}

// LocalRegisterDTO -  회원가입시 필요한 데이터를 객체화
type LocalRegisterDTO struct {
	Email       string `json:"email"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	ShortBio    string `json:"short_bio"`
	UserID      string `json:"user_id"`
}

// SocialRegisterDTO - 소셜 회원가입에 필요한 데이터
type SocialRegisterDTO struct {
	DisplayName string `json:"display_name" binding:"required"`
	UserName    string `json:"username" binding:"required"`
	ShortBio    string `json:"short_bio"`
}

// RegisterTokenDTO - registerToken 내에 정의된 데이터 정보
type RegisterTokenDTO struct {
	Profile social.SocialProfile `json:"profile"`
	Token string `json:"token"`
	Provider string `json:"provider"`
}
