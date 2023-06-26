package controllers

import (
	"BNMO/database"
	"BNMO/models"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Admin accept or reject requests
func ValidateRequests(c *gin.Context) {
	var validate models.ValidateRequest
	var account models.Account
	var request models.Request

	// Bind arriving json into a validate request model
	err := c.BindJSON(&validate)
	if err != nil {
		log.Println("Validate request failed: Failed binding json", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Failed binding request"})
		return
	}

	database.DB.Find(&request, validate.RequestID)

	// If status is accepted, start procedures
	if validate.Status == "accepted" {
		// Check statements
		// Pull data from request and account tables
		database.DB.Find(&account, request.DestinationID)

		// Request type: add
		if request.RequestType == "add" {
			newBalance := account.Balance + request.ConvertedAmount
			database.DB.Model(&account).Update("balance", newBalance)
		}

		// Request type: subtract
		if request.RequestType == "subtract" {
			// If balance is insufficient, reject the request
			if account.Balance < request.ConvertedAmount {
				database.DB.Model(&request).Update("status", "rejected")
				c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
				return
			} else {
				newBalance := account.Balance - request.ConvertedAmount
				database.DB.Model(&account).Update("balance", newBalance)
			}
		}

		// Update value inside request table
		database.DB.Model(&request).Update("status", validate.Status)
		c.JSON(http.StatusOK, gin.H{"message": "Successfully accepted"})
		return

	} else if validate.Status == "rejected" {
		// Update value inside request table
		database.DB.Model(&request).Update("status", validate.Status)
		c.JSON(http.StatusOK, gin.H{"message": "Successfully rejected"})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validate request failed: Unable to parse status"})
		return
	}

}

// Admin display requests
func GetPendingRequests(c *gin.Context) {
	// Specify limitations
	page, _ := strconv.Atoi(c.Query("page"))
	// filterType := c.DefaultQuery("filter", "")
	// filterKey := c.DefaultQuery("key", "")
	limit := 5
	offset := (page - 1) * limit

	var total int64
	var requests []models.Request
	formattedRequests := make([]map[string]interface{}, 0)

	// Pull data from the requests table inside the database
	// Pull only based on the number of offsets and limits specified
	database.DB.Preload("Destination").Where("status=?", "pending").Offset(offset).Limit(limit).Find(&requests)
	database.DB.Model(&models.Request{}).Where("status=?", "pending").Count(&total)

	for _, request := range requests {
		formattedRequests = append(formattedRequests, gin.H{
			"ID":               request.ID,
			"account_number":   request.Destination.AccountNumber,
			"request_type":     request.RequestType,
			"currency":         request.Currency,
			"amount":           request.Amount,
			"converted_amount": request.ConvertedAmount,
		})
	}

	// Return data to frontend
	c.JSON(http.StatusOK, gin.H{
		"data": formattedRequests,
		"metadata": gin.H{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(total) / float64(limit)),
		},
	})
}
