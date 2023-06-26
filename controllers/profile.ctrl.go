package controllers

import (
	"BNMO/database"
	"BNMO/enum"
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

	err := database.DB.Preload("Account").Preload("Address").Where("account_id = ?", id).First(&customer).Error
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
		Country:            customer.Address.Country,
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func EditProfile(c *gin.Context) {
	id := c.Param("id")
	var request models.EditProfileReq
	var customer gormmodels.Customer
	var filePath string

	err := c.Bind(&request)
	if err != nil {
		utils.HandleInternalServerError(c, err, "Edit profile", "Failed to bind request")
		return
	}

	// User changing their image
	if request.ProfilePicture.Size != 0 {
		database.DB.Where("account_id = ?", id).First(&customer)
		if len(customer.ProfilePicturePath) != 0 {
			err := utils.DeleteFile(customer.ProfilePicturePath)
			if err != nil {
				utils.HandleInternalServerError(c, err, "Edit profile", "Failed to delete file")
				return
			}
		}
		filePath = utils.SaveFile(c, request.ProfilePicture, enum.FILE_PROFILE_PICTURE)
	}

	database.DB.Preload("Account").Preload("Address").Where("account_id = ?", id).Updates(gormmodels.Customer{
		Account: gormmodels.Account{
			FirstName: request.FirstName,
			LastName:  request.LastName,
		},
		PhoneNumber:        request.PhoneNumber,
		ProfilePicturePath: filePath,
		Address: gormmodels.CustomerAddress{
			AddressLine1: request.AddressLine1,
			AddressLine2: request.AddressLine2,
			City:         request.City,
			State:        request.State,
			PostalCode:   request.PostalCode,
			Country:      request.Country,
		},
	}).First(&customer)

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
		Country:            customer.Address.Country,
	}

	c.JSON(http.StatusOK, gin.H{"data": response, "message": "Edit successful"})
}
