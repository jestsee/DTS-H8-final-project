package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string
	Email       string
	Password    string
	Age         uint
	Photos      []Photo     `gorm:"foreignKey:User_id"`
	Comments    []Comment   `gorm:"foreignKey:User_id"`
	SocialMedia SocialMedia `gorm:"foreignKey:User_id"`
}
