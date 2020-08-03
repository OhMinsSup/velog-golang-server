package dto

type UpdateProfileBody struct {
	DisplayName string `json:"display_name"`
	ShortBio    string `json:"short_bio"`
	Thumbnail   string `json:"thumbnail"`
}

type UpdateEmailRulesBody struct {
	EmailNotification bool `json:"email_notification"`
	EmailPromotion    bool `json:"email_promotion"`
}

type UpdateSocialInfoBody struct {
	Twitter  string `json:"twitter"`
	Facebook string `json:"facebook"`
	Github   string `json:"github"`
}

