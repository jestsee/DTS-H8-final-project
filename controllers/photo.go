package controllers

import (
	"myGram/models"
	"myGram/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPhotos godoc
// @Summary Get all photos of speicifc user
// @Description Get all photos of speicifc user
// @Tag photo
// @Produce json
// @Success 200 {array} models.GetPhotoResponse
// @Router /photos [get]
func (idb *InDB) GetPhotos(c *gin.Context) {
	var (
		photos []models.Photo
		user   models.User
	)
	userId := utils.GetUserId(c)

	err := idb.DB.Find(&photos, "user_id = ?", userId).Error
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

	for i := range photos {
		photos[i].User.Email = user.Email
		photos[i].User.Username = user.Username
	}

	c.JSON(http.StatusOK, photos)
}

// AddPhoto godoc
// @Summary Add new photo
// @Description Add new photo
// @Tag photo
// @Produce json
// @Param user body models.CreatePhotoRequest true "Create photo"
// @Success 201 {object} models.CreatePhotoResponse
// @Router /photos [post]
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

	resp := map[string]interface{}{
		"id":         photo.Id,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.Photo_url,
		"user_id":    photo.User_id,
		"created_at": photo.CreatedAt,
	}

	c.JSON(http.StatusCreated, resp)
}

// Update godoc
// @Summary Update existing photo
// @Description Update existing photo
// @Tag photo
// @Produce json
// @Security Bearer
// @Param authorization header string true "Authorization"
// @Param photoId query int true "Update photo"
// @Param photo body models.UpdatePhotoRequest true "Update photo"
// @Success 200 {object} models.UpdatePhotoResponse
// @Router /photos [put]
func (idb *InDB) UpdatePhoto(c *gin.Context) {
	var photo *models.Photo
	photoId := c.Param("photoId")
	userId := utils.GetUserId(c)

	err := idb.DB.First(&photo, photoId).Error
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

	c.Bind(&photo)

	err = idb.DB.Save(&photo).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	resp := map[string]interface{}{
		"id":         photo.Id,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.Photo_url,
		"user_id":    photo.User_id,
		"updated_at": photo.UpdatedAt,
	}
	c.JSON(http.StatusOK, resp)
}

// Delete godoc
// @Summary Delete existing photo
// @Description Delete existing photo
// @Tag Photo
// @Produce json
// @Security Bearer
// @Param authorization header string true "Authorization"
// @Param photoId query int true "Delete photo"
// @Success 200 {object} models.DeleteResponse
// @Router /photos [delete]
func (idb *InDB) DeletePhoto(c *gin.Context) {
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
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
		return
	}

	c.Bind(&photo)

	err = idb.DB.Delete(&photo).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
