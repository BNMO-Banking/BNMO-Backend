package controllers

import (
	"BNMO/database"
	gormmodels "BNMO/gorm_models"
	"BNMO/models"
	"BNMO/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Customer add and subtract requests
func AddRequest(c *gin.Context) {
	var request models.RequestReq
	var customer gormmodels.Customer

	err := c.BindJSON(&request)
	if err != nil {
		utils.HandleInternalServerError(c, err, "Add request", "Failed to bind request")
		return
	}

	database.DB.Where("id = ?", request.Id).First(&customer)
	// Check pin validity
	combined := utils.CombinePin(request.Id, request.Pin)
	err = utils.ComparePin(customer.Pin, combined)
	if err != nil {
		utils.HandleBadRequest(c, "Incorrect PIN")
		return
	}

	// Calculate conversion rate
	converted_amount := calculateConversion(request.Currency, request.Amount)

	database.DB.Create(&gormmodels.Request{
		RequestType:     request.RequestType,
		Currency:        request.Currency,
		Amount:          request.Amount,
		ConvertedAmount: converted_amount,
		CustomerID:      request.Id,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Request successfully added. Please wait for validation"})
}
