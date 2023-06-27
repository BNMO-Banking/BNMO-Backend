package controllers

import (
	"BNMO/database"
	gormmodels "BNMO/gorm_models"
	"BNMO/models"
	"BNMO/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetPin(c *gin.Context) {
	var request models.PinReq

	err := c.BindJSON(&request)
	if err != nil {
		utils.HandleInternalServerError(c, err, "Set pin", "Failed to bind request")
		return
	}

	// Hashing pin
	pin, err := utils.HashPin(request.Id, request.Pin)
	if err != nil {
		utils.HandleInternalServerError(c, err, "Set pin", "Failed to hash pin")
		return
	}

	database.DB.Where("id = ?", request.Id).Updates(gormmodels.Customer{
		Pin: pin,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Successfully set pin"})
}

func ComparePin(c *gin.Context) {
	var request models.PinReq
	var customer gormmodels.Customer

	err := c.BindJSON(&request)
	if err != nil {
		utils.HandleInternalServerError(c, err, "Compare pin", "Failed to bind request")
		return
	}

	database.DB.Where("id = ?", request.Id).First(&customer)
	combined := utils.CombinePin(request.Id, request.Pin)
	err = utils.ComparePin(customer.Pin, combined)
	if err != nil {
		utils.HandleBadRequest(c, "Incorrect PIN")
		return
	}

	c.JSON(http.StatusOK, gin.H{"valid": true})
}
