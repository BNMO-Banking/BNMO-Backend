package controllers

import (
	"BNMO/database"
	gormmodels "BNMO/gorm_models"
	"BNMO/models"
	"BNMO/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddAssociates(c *gin.Context) {
	var request models.AddAssociatesReq
	var sourceCustomer gormmodels.Customer
	var destinationCustomer gormmodels.Customer

	err := c.BindJSON(&request)
	if err != nil {
		utils.HandleInternalServerError(c, err, "Add associates", "Failed to bind request")
		return
	}

	database.DB.Preload("Associates").Where("id = ?", request.Id).First(&sourceCustomer)

	err = database.DB.Where("account_number = ?", request.DestinationNumber).First(&destinationCustomer).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.HandleBadRequest(c, "Destination not found")
		return
	}

	for _, associate := range sourceCustomer.Associates {
		if associate.ID == destinationCustomer.ID {
			utils.HandleBadRequest(c, "Account already added")
			return
		}
	}

	database.DB.Model(&sourceCustomer).Omit("Associates.*").Association("Associates").Append(&destinationCustomer)
	c.JSON(http.StatusOK, gin.H{"message": "Destination successfully added"})
}

func CheckAssociates(c *gin.Context) {
	destinationNumber := c.Param("number")
	var customer gormmodels.Customer

	err := database.DB.Preload("Account").Where("account_number = ?", destinationNumber).First(&customer).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.HandleBadRequest(c, "Destination not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"name": customer.Account.FirstName + " " + customer.Account.LastName})
}

func GetAssociates(c *gin.Context) {
	var response []models.DestinationsRes

	id := c.Query("id")

	database.DB.
		Model(&gormmodels.Customer{}).
		Select("associates.account_number, accounts.first_name, accounts.last_name").
		Joins("JOIN customer_associates ON customer_associates.customer_id = customers.id").
		Joins("JOIN customers associates ON associates.id = customer_associates.associate_id").
		Joins("JOIN accounts ON accounts.id = associates.account_id").
		Where("customers.id = ?", id).
		First(&response)

	c.JSON(http.StatusOK, models.DestinationResList{
		Data: response,
	})
}
