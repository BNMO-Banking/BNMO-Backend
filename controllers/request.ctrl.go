package controllers

import (
	"BNMO/database"
	"BNMO/models"
	"log"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Customer add and subtract requests
func AddRequest(c *gin.Context) {
	var request models.Request

	// Bind arriving json into a request model
	err := c.BindJSON(&request)
	if err != nil {
		log.Println("Adding request failed: Failed binding json", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Failed binding request"})
		return
	}

	// Calculate conversion rate
	_, rate := getRatesFromRedis(request.Currency)
	conversion := float64(request.Amount) / rate
	newAmount := int64(math.Floor(conversion))
	request.ConvertedAmount = newAmount

	create := database.DATABASE.Create(&request)
	if create.Error != nil {
		log.Println("Adding request failed: Failed inserting to database", create.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Failed inserting request to database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Request successfully added. Please wait for validation"})
}