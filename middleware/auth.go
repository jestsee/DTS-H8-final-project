package middleware

import (
	"myGram/config"
	"myGram/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication(conf config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := utils.VerifyToken(c, conf)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H {
				"error" : "Unauthenticated",
				"message": err.Error(),
			})
			return
		}
		c.Set("userData", verifyToken)
		c.Next()
	}
}