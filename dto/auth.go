package dto

type SendEmailBody struct {
	Email string `json:"email", binding:"exists,email,required"`
}

type LocalRegisterBody struct {
	RegisterToken string `json:"register_token", binding:"required"`
	DisplayName   string `json:"display_name", binding:"required"`
	UserName      string `json:"username", binding:"required"`
	ShortBio      string `json:"short_bio"`
}

type RegisterTokenPayload struct {
	Email string `json:"email"`
	ID    string `json:"id"`
}

type RegisterTokenType struct {
	Exp     int64 `json:"exp"`
	Issuer  string `json:"issuer"`
	Payload RegisterTokenPayload `json:"payload"`
	Subject string `json:"subject"`
}
