package controllers

import (
	"BNMO/database"
	"BNMO/enum"
	gormmodels "BNMO/gorm_models"
	"BNMO/models"
	"BNMO/utils"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func GetPendingRequests(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit := 10
	offset := (page - 1) * limit

	var total int64
	var requests []models.RequestData

	database.DB.
		Model(&gormmodels.Request{}).
		Select("requests.id, requests.request_type, requests.currency, requests.amount, requests.converted_amount, requests.status, requests.remarks, accounts.first_name, accounts.last_name, customers.account_number, customers.phone_number").
		Joins("JOIN customers ON requests.customer_id = customers.id").
		Joins("JOIN accounts ON customers.account_id = accounts.id").
		Scan(&requests).
		Offset(offset).
		Limit(limit).
		Count(&total)

	// Return data to frontend
	c.JSON(http.StatusOK, models.RequestDataList{
		Data: requests,
		Metadata: models.PageMetadata{
			Total:    total,
			Page:     page,
			LastPage: math.Ceil(float64(total) / float64(limit)),
		},
	})
}

func ValidateRequest(c *gin.Context) {
	id := c.Param("id")
	status := c.Param("status")

	// If status is accepted, start procedures
	if status == string(enum.REQUEST_ACCEPTED) {
		var request gormmodels.Request
		var newBalance decimal.Decimal
		database.DB.Preload("Customer").Where("id = ?", id).First(&request)

		if request.RequestType == enum.ADD {
			newBalance = request.Customer.Balance.Add(request.ConvertedAmount)
		}

		// Request type: subtract
		if request.RequestType == enum.SUBTRACT {
			// If balance is insufficient, reject the request
			if request.Customer.Balance.LessThan(request.ConvertedAmount) {
				database.DB.Where("id = ?", id).Updates(gormmodels.Request{
					Status: enum.REQUEST_REJECTED,
				})
				utils.HandleBadRequest(c, "Insufficient funds")
				return
			} else {
				newBalance = request.Customer.Balance.Sub(request.ConvertedAmount)
			}
		}

		database.DB.Where("id = ?", id).Updates(gormmodels.Request{
			Status: enum.REQUEST_ACCEPTED,
		})

		database.DB.Where("id = ?", request.CustomerID).Updates(gormmodels.Customer{
			Balance: newBalance,
		})
		c.JSON(http.StatusOK, gin.H{"message": "Request successfully accepted"})
		return
	} else if status == string(enum.REQUEST_REJECTED) {
		database.DB.Where("id = ?", id).Updates(gormmodels.Request{
			Status: enum.REQUEST_REJECTED,
		})
		c.JSON(http.StatusOK, gin.H{"message": "Request successfully rejected"})
		return
	}
}
