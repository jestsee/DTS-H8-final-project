package models

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title      string
	Caption    string
	Photo_url  string
	User_id    uint
	Comments    []Comment   `gorm:"foreignKey:Photo_id"`
}
