package controllers

import (
	"myGram/models"
	"myGram/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetComments(c *gin.Context) {
	var comments []models.Comment
	userId := utils.GetUserId(c)

	err := idb.DB.Find(&comments, "user_id = ?", userId).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, comments)
}

// TODO harus autentikasi dan autorisasi
// berarti cuma boleh nambahin komen di foto sendiri kah?
func (idb *InDB) AddComment(c *gin.Context) {
	var (
		comment models.Comment
		photo   models.Photo
	)
	userId := utils.GetUserId(c)

	c.Bind(&comment)

	err := idb.DB.First(&photo, comment.Photo_id).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if photo.User_id != userId {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
		return
	}

	comment.User_id = userId

	err = idb.DB.Debug().Create(&comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	data := map[string]interface{}{
		"id":         comment.Id,
		"message":    comment.Message,
		"photo_id":   comment.Photo_id,
		"user_id":    comment.User_id,
		"created_at": comment.CreatedAt,
	}

	c.JSON(http.StatusCreated, data)
}
