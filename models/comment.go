package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	User_id  uint   `json:"user_id"`
	Photo_id uint   `json:"photo_id"`
	Message  string `json:"message"`
}
