package middlewares

import (
	"net/http"

	"github.com/botsgalaxy/Event-Booking/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized",
		})
		return

	}
	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized by verify token",
		})
		return

	}
	c.Set("userId", userId)
	c.Next()
}
