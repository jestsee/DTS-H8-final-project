package models

type SocialMedia struct {
	Base
	Name             string `json:"name"`
	Social_media_url string `json:"social_media_url"`
	User_id          uint   `json:"user_id"`
}
