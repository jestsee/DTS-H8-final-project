package controllers

import (
	"myGram/models"
	"myGram/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSocialMedias godoc
// @Summary Get all social medias of speicifc user
// @Description Get all social medias of speicifc user
// @Tag socialMedia
// @Produce json
// @Success 200 {object} models.GetSocialMediaResponse
// @Router /socialMedias [get]
func (idb *InDB) GetSocialMedias(c *gin.Context) {
	var (
		socials []models.SocialMedia
		user    models.User
	)
	userId := utils.GetUserId(c)

	err := idb.DB.Find(&socials, "user_id = ?", userId).Error
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

	for i := range socials {
		socials[i].User.Id = user.Id
		socials[i].User.Username = user.Username
	}

	c.JSON(http.StatusOK, gin.H{
		"social_medias": socials,
	})
}

// AddSocialMedia godoc
// @Summary Add new social media
// @Description Add new social media
// @Tag socialMedia
// @Produce json
// @Param user body models.CreateSocialMediaRequest true "Create social media"
// @Success 201 {object} models.CreateSocialMediaResponse
// @Router /socialMedias [post]
func (idb *InDB) AddSocialMedia(c *gin.Context) {
	var social models.SocialMedia
	userId := utils.GetUserId(c)

	c.Bind(&social)

	social.User_id = userId

	err := idb.DB.Debug().Create(&social).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	data := map[string]interface{}{
		"id":               social.Id,
		"name":             social.Name,
		"social_media_url": social.Social_media_url,
		"user_id":          social.User_id,
		"created_at":       social.CreatedAt,
	}

	c.JSON(http.StatusCreated, data)
}

// Update godoc
// @Summary Update existing social media
// @Description Update existing social media
// @Tag socialMedia
// @Produce json
// @Security Bearer
// @Param authorization header string true "Authorization"
// @Param socialMediaId query int true "Update social media"
// @Param socialMedia body models.UpdateSocialMediaRequest true "Update social media"
// @Success 200 {object} models.UpdateSocialMediaResponse
// @Router /socialMedias [put]
func (idb *InDB) UpdateSocialMedia(c *gin.Context) {
	var socials *models.SocialMedia
	socialMediaId := c.Param("socialMediaId")
	userId := utils.GetUserId(c)

	err := idb.DB.First(&socials, socialMediaId).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if socials.User_id != userId {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
		return
	}

	c.Bind(&socials)

	err = idb.DB.Save(&socials).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	resp := map[string]interface{}{
		"id":               socials.Id,
		"name":             socials.Name,
		"social_media_url": socials.Social_media_url,
		"user_id":          socials.User_id,
		"updated_at":       socials.UpdatedAt,
	}
	c.JSON(http.StatusOK, resp)
}

// Delete godoc
// @Summary Delete existing social media
// @Description Delete existing social media
// @Tag socialMedia
// @Produce json
// @Security Bearer
// @Param authorization header string true "Authorization"
// @Param socialMediaId query int true "Delete social media"
// @Success 200 {object} models.DeleteResponse
// @Router /socialMedias [delete]
func (idb *InDB) DeleteSocialMedia(c *gin.Context) {
	var social *models.SocialMedia
	socialMediaId := c.Param("socialMediaId")
	userId := utils.GetUserId(c)

	err := idb.DB.First(&social, socialMediaId).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if social.User_id != userId {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
		return
	}

	c.Bind(&social)

	err = idb.DB.Delete(&social).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
