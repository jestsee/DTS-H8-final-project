package models

type Comment struct {
	Base
	User_id  uint   `json:"user_id"`
	Photo_id uint   `json:"photo_id"`
	Message  string `json:"message" gorm:"not null;default:null"`
}
