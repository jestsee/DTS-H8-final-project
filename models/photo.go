package models

type Photo struct {
	Base
	Title     string `json:"title" gorm:"not null;default:null"`
	Caption   string `json:"caption"`
	Photo_url string `json:"photo_url" gorm:"not null;default:null"`
	User_id   uint   `json:"user_id"`
	
	Comments []Comment `json:"-" gorm:"foreignKey:Photo_id"`
	
	User      UserP  `gorm:"-"`
}

type UserP struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
