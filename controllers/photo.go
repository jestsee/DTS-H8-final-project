package controllers

import (
	"myGram/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) AddPhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := GetContentType(c)

	photo := models.Photo{}
	userId := uint(userData["id"].(float64))

	if contentType == "application/json" {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

	photo.User_id = userId

	err := idb.DB.Debug().Create(&photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, photo)
}
