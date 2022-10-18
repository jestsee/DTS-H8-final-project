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

func (u *User) BeforeDelete(tx *gorm.DB) error {
	err := tx.Model(u.SocialMedia).Where("user_id = ?", u.Id).Delete(u.SocialMedia).Error
	if err != nil {
		return err
	}
	err = tx.Model(u.Comments).Where("user_id = ?", u.Id).Delete(u.Comments).Error
	if err != nil {
		return err
	}
	err = tx.Model(u.Photos).Where("user_id = ?", u.Id).Delete(u.Photos).Error
	return err
}

func (p *Photo) BeforeDelete(tx *gorm.DB) error {
	err := tx.Model(p.Comments).Where("photo_id = ?", p.Id).Delete(p.Comments).Error
	if err != nil {
		return err
	}
	return nil
}
