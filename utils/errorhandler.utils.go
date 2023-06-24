package utils

import (
	"BNMO/enum"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error, error_type enum.ErrorType, function string, message string) {
	errorMessage := fmt.Sprintf("%s failed: %s", function, message)
	if error_type == enum.BAD_REQUEST {
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
	} else if error_type == enum.INTERNAL_SERVER_ERROR {
		log.Println(fmt.Sprintf("%s with error %s", errorMessage, err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": errorMessage})
	}
	return
}
