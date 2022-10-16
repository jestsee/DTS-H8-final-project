package models

import (
	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model
	Name             string
	Social_media_url string
	User_id           uint
}
