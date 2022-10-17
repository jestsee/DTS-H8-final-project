package controllers

import (
	"myGram/models"
	"myGram/utils"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func (idb *InDB) Register(c *gin.Context) {
	var (
		user models.User
		data map[string]interface{}
		err  error
	)
	c.Bind(&user)

	err = idb.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data = map[string]interface{}{
		"age":      user.Age,
		"email":    user.Email,
		"id":       user.Id,
		"username": user.Username,
	}

	c.JSON(http.StatusCreated, gin.H(data))
}

func (idb *InDB) Login(c *gin.Context) {
	var (
		login models.Login
		user  models.User
	)

	c.Bind(&login)

	// password checking
	err := idb.DB.First(&user, "email = ?", login.Email).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	valid := utils.CheckPasswordHash(login.Password, user.Password)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "password does not match",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": utils.GenerateToken(user.Id, user.Email, *idb.Conf),
	})
}

func (idb *InDB) UpdateUser(c *gin.Context) {
	var user models.User
	userId := utils.GetUserId(c)
	user.Id = userId

	c.Bind(&user)

	err := idb.DB.Model(&user).Clauses(clause.Returning{}).Updates(models.User{Username: user.Username, Email: user.Email}).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	data := map[string]interface{}{
		"id":         user.Id,
		"email":      user.Email,
		"username":   user.Username,
		"age":        user.Age,
		"updated_at": user.UpdatedAt,
	}
	c.JSON(http.StatusOK, data)
}

func (idb *InDB) DeleteUser(c *gin.Context) {
	var user models.User
	userId := utils.GetUserId(c)
	user.Id = userId

	err := idb.DB.Delete(&user).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}