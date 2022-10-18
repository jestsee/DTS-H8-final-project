package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"myGram/models"
	"myGram/utils"
	"net/http"
)

// Register godoc
// @Summary Register new user
// @Description Register new user
// @Tag User
// @Produce json
// @Param user body models.RegisterRequest true "Create user"
// @Success 201 {object} models.RegisterResponse
// @Router /users/register [post]
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

// Login godoc
// @Summary Login existing user
// @Description Login existing user
// @Tag User
// @Produce json
// @Param user body models.LoginRequest true "Login user"
// @Success 201 {object} models.LoginResponse
// @Router /users/login [post]
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

// Update godoc
// @Summary Update existing user
// @Description Update existing user
// @Tag User
// @Produce json
// @Security Bearer
// @Param authorization header string true "Authorization"
// @Param userId query int true "Update user"
// @Param user body models.UpdateUserRequest true "Update user"
// @Success 200 {object} models.UpdateUserResponse
// @Router /users [put]
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

// Delete godoc
// @Summary Delete existing user
// @Description Delete existing user
// @Tag User
// @Produce json
// @Security Bearer
// @Param authorization header string true "Authorization"
// @Param userId query int true "Delete user"
// @Success 200 {object} models.DeleteResponse
// @Router /users [delete]
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
