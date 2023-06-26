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

func GetProfile(c *gin.Context) {
	id := c.Param("id")
	var customer gormmodels.Customer

	err := database.DB.Preload("Account").Preload("Address").Where("account_id=?", id).First(&customer).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.HandleRecordNotFound(c, "Get profile", "")
		return
	}

	response := models.ProfileRes{
		AccountNumber:      customer.AccountNumber,
		AccountType:        customer.Account.AccountType,
		Email:              customer.Account.Email,
		Username:           customer.Account.Username,
		FirstName:          customer.Account.FirstName,
		LastName:           customer.Account.LastName,
		CardNumber:         customer.CardNumber,
		Balance:            customer.Balance,
		PhoneNumber:        customer.PhoneNumber,
		ProfilePicturePath: customer.ProfilePicturePath,
		AddressLine1:       customer.Address.AddressLine1,
		AddressLine2:       customer.Address.AddressLine2,
		City:               customer.Address.City,
		State:              customer.Address.State,
		PostalCode:         customer.Address.PostalCode,
		Country:            customer.Address.PostalCode,
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}
