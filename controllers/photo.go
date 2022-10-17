package controllers

import (
	"myGram/models"
	"myGram/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetPhotos(c *gin.Context) {
	var photos []models.Photo
	userId := utils.GetUserId(c)

	err := idb.DB.Find(&photos, "User_id = ?", userId).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, photos)
}

func (idb *InDB) AddPhoto(c *gin.Context) {
	var photo models.Photo
	userId := utils.GetUserId(c)

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

func (idb *InDB) UpdatePhoto(c *gin.Context) {
	var photo *models.Photo
	id := c.Param("photoId")
	userId := utils.GetUserId(c)

	err := idb.DB.First(&photo, id).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if photo.User_id != userId {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{
				"error": "Unauthorized", 
				"message": "you are not allowed to access this data",
			})
		return
	}

	c.Bind(&photo)

	err = idb.DB.Save(&photo).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, photo)
}
