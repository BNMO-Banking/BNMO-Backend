package controllers

import (
	"BNMO/database"
	"BNMO/enum"
	gormmodels "BNMO/gorm_models"
	"BNMO/models"
	"BNMO/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Transfer(c *gin.Context) {
	var request models.TransferReq
	var sourceCustomer gormmodels.Customer
	var destinationCustomer gormmodels.Customer

	err := c.BindJSON(&request)
	if err != nil {
		utils.HandleInternalServerError(c, err, "Transfer", "Failed to bind request")
		return
	}

	database.DB.Where("id = ?", request.Id).First(&sourceCustomer)
	// Check pin validity
	combined := utils.CombinePin(request.Id, request.Pin)
	err = utils.ComparePin(sourceCustomer.Pin, combined)
	if err != nil {
		utils.HandleBadRequest(c, "Incorrect PIN")
		return
	}

	database.DB.Where("account_number = ?", request.DestinationNumber).First(&destinationCustomer)
	// Calculate conversion rate
	converted_amount := calculateConversion(request.Currency, request.Amount)

	if sourceCustomer.Balance.LessThan(converted_amount) {
		database.DB.Create(&gormmodels.Transfer{
			Currency:        request.Currency,
			Amount:          request.Amount,
			ConvertedAmount: converted_amount,
			Status:          enum.TRANSFER_FAILED,
			Description:     request.Description,
			SourceID:        sourceCustomer.ID,
			DestinationID:   destinationCustomer.ID,
		})
		utils.HandleBadRequest(c, "Insufficient funds")
		return
	}

	newSourceBalance := sourceCustomer.Balance.Sub(converted_amount)
	newDestinationBalance := destinationCustomer.Balance.Add(converted_amount)

	// Update database values
	database.DB.Model(&sourceCustomer).Update("balance", newSourceBalance)
	database.DB.Model(&destinationCustomer).Update("balance", newDestinationBalance)
	database.DB.Create(&gormmodels.Transfer{
		Currency:        request.Currency,
		Amount:          request.Amount,
		ConvertedAmount: converted_amount,
		Status:          enum.TRANSFER_SUCCESS,
		Description:     request.Description,
		SourceID:        sourceCustomer.ID,
		DestinationID:   destinationCustomer.ID,
	})
	c.JSON(http.StatusOK, gin.H{"message": "Transfer completed"})
}
