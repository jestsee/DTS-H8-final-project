package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string      `json:"username"`
	Email       string      `json:"email"`
	Password    string      `json:"-"`
	Age         uint        `json:"age"`
	Photos      []Photo     `gorm:"foreignKey:User_id"`
	Comments    []Comment   `gorm:"foreignKey:User_id"`
	SocialMedia SocialMedia `gorm:"foreignKey:User_id"`
}
