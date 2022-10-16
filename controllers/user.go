package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "hello world!",
	})
}