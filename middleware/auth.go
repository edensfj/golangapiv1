package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			c.JSON(http.StatusForbidden, gin.H{"code": http.StatusForbidden, "message": "Missing Auth Token"})
			c.Abort()
			return
		}
		token := strings.Split(tokenHeader, " ")
		if len(token) != 2 {
			c.JSON(http.StatusForbidden, gin.H{"code": http.StatusForbidden, "message": "Invalid auth token"})
			c.Abort()
			return
		}
		c.Next()
	}
}
