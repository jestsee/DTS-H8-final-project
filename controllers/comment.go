package controllers

import (
	"myGram/models"
	"myGram/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetComments godoc
// @Summary Get all comments of speicifc user
// @Description Get all comments of speicifc user
// @Tag comment
// @Produce json
// @Success 200 {array} models.GetCommentResponse
// @Router /comments [get]
func (idb *InDB) GetComments(c *gin.Context) {
	var (
		comments []models.Comment
		user     models.User
	)
	userId := utils.GetUserId(c)

	err := idb.DB.Find(&comments, "user_id = ?", userId).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	err = idb.DB.First(&user, userId).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	for i := range comments {
		comments[i].User.Id = user.Id
		comments[i].User.Email = user.Email
		comments[i].User.Username = user.Username

		photo := models.Photo{}

		err = idb.DB.Find(&photo, comments[i].Photo_id).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad request",
				"message": err.Error(),
			})
			return
		}

		comments[i].Photo.Id = photo.Id
		comments[i].Photo.Title = photo.Title
		comments[i].Photo.Caption = photo.Caption
		comments[i].Photo.Photo_url = photo.Photo_url
		comments[i].Photo.User_id = photo.User_id
	}

	c.JSON(http.StatusOK, comments)
}

// TODO harus autentikasi dan autorisasi
// berarti cuma boleh nambahin komen di foto sendiri kah?
// AddComment godoc
// @Summary Add new comment
// @Description Add new comment
// @Tag comment
// @Produce json
// @Param user body models.CreateCommentRequest true "Create comment"
// @Success 201 {object} models.CreateCommentResponse
// @Router /comments [post]
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

// Update godoc
// @Summary Update existing comment
// @Description Update existing comment
// @Tag comment
// @Produce json
// @Security Bearer
// @Param authorization header string true "Authorization"
// @Param commentId query int true "Update comment"
// @Param comment body models.UpdateCommentRequest true "Update comment"
// @Success 200 {object} models.UpdateCommentResponse
// @Router /comments [put]
func (idb *InDB) UpdateComment(c *gin.Context) {
	var comment *models.Comment
	commentId := c.Param("commentId")
	userId := utils.GetUserId(c)

	err := idb.DB.First(&comment, commentId).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if comment.User_id != userId {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
		return
	}

	c.Bind(&comment)

	err = idb.DB.Save(&comment).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// TODO salah ya outputnya?
	resp := map[string]interface{}{
		"id":         comment.Id,
		"user_id":    comment.User_id,
		"photo_id":   comment.Photo_id,
		"message":    comment.Message,
		"updated_at": comment.UpdatedAt,
	}
	c.JSON(http.StatusOK, resp)
}

// Delete godoc
// @Summary Delete existing comment
// @Description Delete existing comment
// @Tag Comment
// @Produce json
// @Security Bearer
// @Param authorization header string true "Authorization"
// @Param commentId query int true "Delete comment"
// @Success 200 {object} models.DeleteResponse
// @Router /comments [delete]
func (idb *InDB) DeleteComment(c *gin.Context) {
	var comment *models.Comment
	commentId := c.Param("commentId")
	userId := utils.GetUserId(c)

	err := idb.DB.First(&comment, commentId).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if comment.User_id != userId {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
		return
	}

	c.Bind(&comment)

	err = idb.DB.Delete(&comment).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
