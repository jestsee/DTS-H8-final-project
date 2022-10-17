package controllers

import (
	"myGram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetPhotos(c *gin.Context) {
	var photos []models.Photo
	userId := GetUserId(c)

	err := idb.DB.Find(&photos, "User_id = ?", userId).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error" : "Bad request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, photos)
}

func (idb *InDB) AddPhoto(c *gin.Context) {
	var photo models.Photo
	userId := GetUserId(c)

	c.Bind(&photo)

	photo.User_id = userId

	err := idb.DB.Debug().Create(&photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, photo)
}
