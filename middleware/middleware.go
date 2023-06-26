package middleware

import (
	"BNMO/enum"
	"BNMO/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CustomerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := token.ParseJWT(c)
		if err != nil || strings.Compare(claims["role"].(string), string(enum.CUSTOMER)) != 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := token.ParseJWT(c)
		if err != nil || strings.Compare(claims["role"].(string), string(enum.ADMIN)) != 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
