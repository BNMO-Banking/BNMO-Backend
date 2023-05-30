package middleware

import (
	"BNMO/utilities"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsAuthenticate(c *gin.Context) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		fmt.Println("Cannot get cookie from request", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Could not retrieve cookies"})
		return
	}

	_, err = utilities.ParseJWT(cookie)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthenticated"})
		return
	}

	c.Next()
}