package utils

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleBadRequest(c *gin.Context, function string, message string) {
	errorMessage := fmt.Sprintf("%s failed: %s", function, message)
	c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
	return
}

func HandleInternalServerError(c *gin.Context, err error, function string, message string) {
	errorMessage := fmt.Sprintf("%s failed: %s", function, message)
	log.Println(fmt.Sprintf("%s with error %s", errorMessage, err.Error()))
	c.JSON(http.StatusInternalServerError, gin.H{"error": errorMessage})
	return
}

func HandleRecordNotFound(c *gin.Context, function string, message string) {
	var errorMessage string
	if len(message) > 0 {
		errorMessage = fmt.Sprintf("%s failed: %s", function, message)
	} else {
		errorMessage = fmt.Sprintf("%s failed: Record not found", function)
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
	return
}
