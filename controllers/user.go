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
	emailValid := IsEmailValid(user.Email)
	if !emailValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email",
		})
		return
	}

	// age validation
	if user.Age <= 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "age must be greater than 8",
		})
		return
	}

	// password validation
	user.Password, err = HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

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
		"id":       user.ID,
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

	// email validation
	valid := IsEmailValid(login.Email)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email",
		})
		return
	}

	// password validation
	valid = IsPasswordValid(login.Password)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "password does not match",
		})
		return
	}

	// password checking
	err := idb.DB.First(&user, "email = ?", login.Email).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	valid = CheckPasswordHash(login.Password, user.Password)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "password does not match",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": "TODO",
	})
}
