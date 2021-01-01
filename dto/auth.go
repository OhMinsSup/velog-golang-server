package dto

// SendEmailBody - SendEmailController 이메일 발송 body 데이터
type SendEmailBody struct {
	Email string `json:"email" binding:"email,required"`
}

// CodeParams - CodeController 코드 인증에 관한 params 데이터
type CodeParams struct {
	Code string `json:"code"`
}

type LocalRegisterBody struct {
	RegisterToken string `json:"register_token" binding:"required"`
	DisplayName   string `json:"display_name" binding:"required"`
	UserName      string `json:"username" binding:"required"`
	ShortBio      string `json:"short_bio"`
}

type SocialRegisterBody struct {
	DisplayName string `json:"display_name" binding:"required"`
	UserName    string `json:"username" binding:"required"`
	ShortBio    string `json:"short_bio"`
}

type RegisterTokenPayload struct {
	Email string `json:"email"`
	ID    string `json:"id"`
}

type RegisterTokenType struct {
	Exp     int64                `json:"exp"`
	Issuer  string               `json:"issuer"`
	Payload RegisterTokenPayload `json:"payload"`
	Subject string               `json:"subject"`
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
