package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateAuthHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("x-is-authentication")
		if authHeader != "xur-2225-vcx-8900-aie" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid Token"})
			c.Abort()
			return
		}
		c.Next()
	}
}
