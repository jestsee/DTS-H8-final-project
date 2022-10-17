package models

type Comment struct {
	Base
	User_id  uint   `json:"user_id"`
	Photo_id uint   `json:"photo_id"`
	Message  string `json:"message" gorm:"not null;default:null"`

	User  UserC  `gorm:"-"`
	Photo PhotoC `gorm:"-"`
}

type UserC struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoC struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_url string `json:"photo_url"`
	User_id   uint   `json:"user_id"`
}
