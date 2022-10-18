package models

type User struct {
	Base
	Username string `json:"username" gorm:"uniqueIndex;not null;default:null"`
	Email    string `json:"email" gorm:"uniqueIndex;not null;default:null"`
	Password string `json:"password" gorm:"not null;default:null"`
	Age      uint   `json:"age" gorm:"not null;default:null"`

	Photos      []Photo       `gorm:"foreignKey:User_id"`
	Comments    []Comment     `gorm:"foreignKey:User_id"`
	SocialMedia []SocialMedia `gorm:"foreignKey:User_id"`
}