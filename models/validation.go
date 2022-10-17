package models

import (
	"errors"
	"myGram/utils"

	"gorm.io/gorm"
)

func (u *User) BeforeCreate(tx *gorm.DB) error {	
	var err error

	// email validation
	valid := utils.IsEmailValid(u.Email)
	if !valid {
		return errors.New(`invalid email`)
	}

	// age validation
	if u.Age <= 8 {
		return errors.New(`age must be greater than 8`)
	}

	// hash password
	u.Password, err = utils.HashPassword(u.Password)
	if err != nil {
		return err 
	}

	return nil
}