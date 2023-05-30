package controller

import (
	"BNMO/database"
	"BNMO/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RequestHistory(c *gin.Context) {
	// Account ID
	id, _ := strconv.Atoi(c.Query("id"))
	// Specify limitations
	page, _ := strconv.Atoi(c.Query("page"))
	limit := 5
	offset := (page-1) * limit

	var total int64
	var requests []models.Request
	var formattedRequests []map[string]interface{}

	// Pull data from the requests table inside the database
	// Pull only based on the number of offsets and limits specified
	database.DATABASE.Where("destination_id=?", id).Offset(offset).Limit(limit).Find(&requests)
	database.DATABASE.Model(&models.Request{}).Where("destination_id=?", id).Count(&total)

	for _, request := range requests {
		formattedRequests = append(formattedRequests, gin.H{
			"ID": request.ID,
			"request_type": request.RequestType,
			"currency": request.Currency,
			"amount": request.Amount,
			"converted_amount": request.ConvertedAmount,
			"status": request.Status,
			"CreatedAt": request.CreatedAt,
		})
	}

	// Return data to frontend
	c.JSON(http.StatusOK, gin.H{
		"data": formattedRequests,
		"metadata": gin.H{
			"total": total,
			"page": page,
			"last_page": math.Ceil(float64(total)/float64(limit)),
		},
	})
}

func TransferHistory(c *gin.Context) {
	// Account ID
	id, _ := strconv.Atoi(c.Query("id"))
	// Specify limitations
	page, _ := strconv.Atoi(c.Query("page"))
	limit := 5
	offset := (page-1) * limit

	var total int64
	var transfers []models.Transfer
	var formattedTransfers []map[string]interface{}

	// Pull data from the requests table inside the database
	// Pull only based on the number of offsets and limits specified
	database.DATABASE.Where("source_id=?", id).Offset(offset).Limit(limit).Find(&transfers)
	database.DATABASE.Model(&models.Transfer{}).Where("source_id=?", id).Count(&total)

	for _, transfer := range transfers {
		formattedTransfers = append(formattedTransfers, gin.H{
			"ID": transfer.ID,
			"destination": transfer.Destination,
			"currency": transfer.Currency,
			"amount": transfer.Amount,
			"converted_amount": transfer.ConvertedAmount,
			"status": transfer.Status,
			"CreatedAt": transfer.CreatedAt,
		})
	}

	// Return data to frontend
	c.JSON(http.StatusOK, gin.H{
		"data": formattedTransfers,
		"metadata": gin.H{
			"total": total,
			"page": page,
			"last_page": math.Ceil(float64(total)/float64(limit)),
		},
	})
}