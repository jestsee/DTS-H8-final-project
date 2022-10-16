package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	User_id    uint
	Photo_id   uint
	Message    string
}