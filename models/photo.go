package models

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title     string    `json:"title" gorm:"not null;default:null"`
	Caption   string    `json:"caption"`
	Photo_url string    `json:"photo_url" gorm:"not null;default:null"`
	User_id   uint      `json:"user_id"`
	Comments  []Comment `gorm:"foreignKey:Photo_id"`
}
