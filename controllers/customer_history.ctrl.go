package controllers

import (
	"BNMO/database"
	gormmodels "BNMO/gorm_models"
	"BNMO/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRequestHistory(c *gin.Context) {
	id := c.Param("id")
	page, _ := strconv.Atoi(c.Query("page"))
	limit := 10
	offset := (page - 1) * limit

	var requests []models.RequestHistory
	var total int64

	database.DB.
		Model(&gormmodels.Request{}).
		Where("customer_id = ?", id).
		Offset(offset).
		Limit(limit).
		Scan(&requests)

	database.DB.
		Model(&gormmodels.Request{}).
		Where("customer_id = ?", id).
		Count(&total)

	c.JSON(http.StatusOK, models.RequestHistoryList{
		Data: requests,
		Metadata: models.PageMetadata{
			Total:    total,
			Page:     page,
			LastPage: math.Ceil(float64(total) / float64(limit)),
		},
	})
}

func GetTransferHistory(c *gin.Context) {
	id := c.Param("id")
	page, _ := strconv.Atoi(c.Query("page"))
	limit := 10
	offset := (page - 1) * limit

	var inboundTransfers []models.TransferHistory
	var outboundTransfers []models.TransferHistory
	var transfers []models.TransferHistory
	var total int64

	database.DB.
		Table("transfers").
		Select("transfers.created_at, transfers.updated_at, transfers.currency, transfers.amount, transfers.converted_amount, transfers.status, transfers.description, customers.account_number, accounts.first_name, accounts.last_name").
		Where("destination_id = ?", id).
		Joins("JOIN customers ON transfers.source_id = customers.id").
		Joins("JOIN accounts ON customers.account_id = accounts.id").
		Scan(&inboundTransfers).
		Offset(offset).
		Limit(limit)

	database.DB.
		Table("transfers").
		Select("transfers.created_at, transfers.updated_at, transfers.currency, transfers.amount, transfers.converted_amount, transfers.status, transfers.description, customers.account_number, accounts.first_name, accounts.last_name").
		Where("source_id = ?", id).
		Joins("JOIN customers ON transfers.destination_id = customers.id").
		Joins("JOIN accounts ON customers.account_id = accounts.id").
		Scan(&outboundTransfers).
		Offset(offset).
		Limit(limit)

	transfers = append(transfers, inboundTransfers...)
	transfers = append(transfers, outboundTransfers...)

	database.DB.
		Model(&gormmodels.Transfer{}).
		Where("source_id = ?", id).
		Count(&total)

	c.JSON(http.StatusOK, models.TransferHistoryList{
		Data: models.TransferTypes{
			InboundTransfers:  inboundTransfers,
			OutboundTransfers: outboundTransfers,
		},
		Metadata: models.PageMetadata{
			Total:    total,
			Page:     page,
			LastPage: math.Ceil(float64(total) / float64(limit)),
		},
	})
}
