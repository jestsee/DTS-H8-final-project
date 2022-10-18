package models

type SocialMedia struct {
	Base
	Name             string `json:"name"`
	Social_media_url string `json:"social_media_url"`
	User_id          uint   `json:"user_id"`

	User UserS `gorm:"-"`
}

type UserS struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
}
