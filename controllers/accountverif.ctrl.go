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
)

func GetPendingAccounts(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit := 10
	offset := (page - 1) * limit

	var total int64
	var accounts []models.AccountData

	database.DB.
		Model(&gormmodels.Customer{}).
		Select("accounts.id, accounts.first_name, accounts.last_name, customers.phone_number, customers.id_card_path, customers.status, customer_addresses.address_line1, customer_addresses.address_line2, customer_addresses.city, customer_addresses.state, customer_addresses.postal_code, customer_addresses.country").
		Joins("JOIN accounts ON customers.account_id = accounts.id").
		Joins("JOIN customer_addresses ON customers.address_id = customer_addresses.id").
		Scan(&accounts).
		Offset(offset).
		Limit(limit).
		Count(&total)

	c.JSON(http.StatusOK, models.AccountDataList{
		Data: accounts,
		Metadata: models.PageMetadata{
			Total:    total,
			Page:     page,
			LastPage: math.Ceil(float64(total) / float64(limit)),
		},
	})
}

func ValidateAccount(c *gin.Context) {
	id := c.Param("id")
	status := c.Param("status")

	if status == string(enum.ACCOUNT_ACCEPTED) {
		database.DB.Where("account_id = ?", id).Updates(gormmodels.Customer{
			Status:        enum.ACCOUNT_ACCEPTED,
			AccountNumber: utils.GenerateAccountNumber(),
			CardNumber:    utils.GenerateCardNumber(),
		})
		c.JSON(http.StatusOK, gin.H{"message": "Account successfully accepted"})
		return
	} else if status == string(enum.ACCOUNT_REJECTED) {
		database.DB.Where("account_id = ?", id).Updates(gormmodels.Customer{
			Status: enum.ACCOUNT_REJECTED,
		})
		c.JSON(http.StatusOK, gin.H{"message": "Account successfully rejected"})
		return
	}

}
