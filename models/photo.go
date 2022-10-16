package models

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Photo_url string    `json:"photo_url"`
	User_id   uint      `json:"user_id"`
	Comments  []Comment `gorm:"foreignKey:Photo_id"`
}
