package controllers

import (
	"myGram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// https://www.sohamkamani.com/golang/jwt-authentication/

func (idb *InDB) Register(c *gin.Context) {
	var (
		user models.User
		data map[string]interface{}
		err  error
	)
	c.Bind(&user)

	// email validation
	_, emailValid := IsEmailValid(user.Email)
	if !emailValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email",
		})
		return
	}

	user.Password, err = HashPassword(user.Password)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	err = idb.DB.Create(&user).Error
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	data = map[string]interface{}{
		"age":      user.Age,
		"email":    user.Email,
		"id":       user.ID,
		"username": user.Username,
	}

	c.JSON(http.StatusCreated, gin.H(data))
}
