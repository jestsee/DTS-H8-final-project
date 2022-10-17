package controllers

import (
	"myGram/config"

	"gorm.io/gorm"
)

type InDB struct {
	DB *gorm.DB
	Conf *config.Config
}