package controllers

import (
	"BNMO/database"
	"BNMO/enum"
	gormmodels "BNMO/gorm_models"
	"BNMO/models"
	"BNMO/utils"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func GetProfile(c *gin.Context) {
	id := c.Param("id")
	var customer gormmodels.Customer

	err := database.DB.Preload("Account").Preload("Address").Where("id = ?", id).First(&customer).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.HandleRecordNotFound(c, "")
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
	if request.ProfilePicture != nil {
		if request.ProfilePicture.Size != 0 {
			database.DB.Where("id = ?", id).First(&customer)
			if len(customer.ProfilePicturePath) != 0 {
				err := utils.DeleteFile(customer.ProfilePicturePath)
				if err != nil {
					utils.HandleInternalServerError(c, err, "Edit profile", "Failed to delete file")
					return
				}
			}
			filePath = utils.SaveFile(c, request.ProfilePicture, enum.FILE_PROFILE_PICTURE)
		}
	}

	database.DB.Preload("Account").Preload("Address").Where("id = ?", id).Updates(&gormmodels.Customer{
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
	}).First((&customer))

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

func GetStatistics(c *gin.Context) {
	id := c.Param("id")
	year := c.Query("year")
	var customer gormmodels.Customer
	var totalReceivedRequest decimal.Decimal
	var totalReceivedTransfer decimal.Decimal

	var totalSpentRequest decimal.Decimal
	var totalSpentTransfer decimal.Decimal

	monthlyReceived := make([]decimal.Decimal, 12)
	monthlySpending := make([]decimal.Decimal, 12)

	err := database.DB.Where("id = ?", id).First(&customer).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.HandleRecordNotFound(c, "")
		return
	}

	database.DB.Table("requests").Select("SUM(converted_amount)").Where("status = ? AND customer_id = ? AND request_type = ?", enum.REQUEST_ACCEPTED, id, enum.ADD).Scan(&totalReceivedRequest)
	database.DB.Table("transfers").Select("SUM(converted_amount)").Where("status = ? AND destination_id = ?", enum.TRANSFER_SUCCESS, id).Scan(&totalReceivedTransfer)

	database.DB.Table("requests").Select("SUM(converted_amount)").Where("status = ? AND customer_id = ? AND request_type = ?", enum.REQUEST_ACCEPTED, id, enum.SUBTRACT).Scan(&totalSpentRequest)
	database.DB.Table("transfers").Select("SUM(converted_amount)").Where("status = ? AND source_id = ?", enum.TRANSFER_SUCCESS, id).Scan(&totalSpentTransfer)

	startDate, _ := time.Parse("2006-01-02", fmt.Sprintf("%s-01-01", year))
	for i := 0; i < 12; i++ {
		var monthlyReceivedRequest decimal.Decimal
		var monthlyReceivedTransfer decimal.Decimal

		var monthlySpentRequest decimal.Decimal
		var monthlySpentTransfer decimal.Decimal
		endDate := startDate.AddDate(0, 1, 0).Add(-time.Nanosecond)

		database.DB.
			Table("requests").
			Select("SUM(converted_amount)").
			Where("status = ? AND customer_id = ? AND request_type = ? AND updated_at >= ? AND updated_at <= ?",
				enum.REQUEST_ACCEPTED, id, enum.ADD, startDate.Format("2006-01-02"), endDate.Format("2006-01-02")).
			Scan(&monthlyReceivedRequest)
		database.DB.
			Table("transfers").
			Select("SUM(converted_amount)").
			Where("status = ? AND destination_id = ? AND updated_at >= ? AND updated_at <= ?",
				enum.TRANSFER_SUCCESS, id, startDate.Format("2006-01-02"), endDate.Format("2006-01-02")).
			Scan(&monthlyReceivedTransfer)

		database.DB.
			Table("requests").
			Select("SUM(converted_amount)").
			Where("status = ? AND customer_id = ? AND request_type = ? AND updated_at >= ? AND updated_at <= ?",
				enum.REQUEST_ACCEPTED, id, enum.SUBTRACT, startDate.Format("2006-01-02"), endDate.Format("2006-01-02")).
			Scan(&monthlySpentRequest)
		database.DB.
			Table("transfers").
			Select("SUM(converted_amount)").
			Where("status = ? AND source_id = ? AND updated_at >= ? AND updated_at <= ?",
				enum.TRANSFER_SUCCESS, id, startDate.Format("2006-01-02"), endDate.Format("2006-01-02")).
			Scan(&monthlySpentTransfer)

		monthlyReceived[i] = monthlyReceivedRequest.Add(monthlyReceivedTransfer)
		monthlySpending[i] = monthlySpentRequest.Add(monthlySpentTransfer)

		startDate = startDate.AddDate(0, 1, 0)
	}

	response := models.ProfileStatistics{
		Balance:         customer.Balance,
		TotalSpent:      totalSpentRequest.Add(totalSpentTransfer),
		TotalReceived:   totalReceivedRequest.Add(totalReceivedTransfer),
		MonthlySpending: monthlySpending,
		MonthlyReceived: monthlyReceived,
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}
