package controller

import (
	"BNMO/database"
	"BNMO/models"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddDestination(c *gin.Context) {
	var request models.AddDestinationRequest
	var userAccount models.Account
	var destinationAccount models.Account

	// Bind arriving json to add destination model
	err := c.BindJSON(&request)
	if err != nil {
		log.Println("Add destination failed: Failed binding json", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Failed binding request"})
		return
	}

	database.DATABASE.Find(&userAccount, request.Id)
	fmt.Println(userAccount.ID, userAccount.FirstName)

	database.DATABASE.Where("account_number=?", request.DestinationNumber).First(&destinationAccount)
	fmt.Println(destinationAccount.ID, destinationAccount.FirstName)
	if destinationAccount.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No account found with given number"})
		return
	}

	userAccount.TransferDestination = append(userAccount.TransferDestination, &destinationAccount)

	database.DATABASE.Save(&userAccount)
	c.JSON(http.StatusOK, gin.H{"message": "Destination successfully added"})
}

func CheckDestination(c *gin.Context) {
	var request models.CheckDestinationRequest
	var account models.Account

	// Bind arriving json to check destination model
	err := c.BindJSON(&request)
	if err != nil {
		log.Println("Check destination failed: Failed binding json", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Failed binding request"})
		return
	}

	database.DATABASE.Where("account_number=?", request.DestinationNumber).First(&account)
	if account.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No account found with given number"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"name": account.FirstName + " " +  account.LastName})
}

func GetDestination(c *gin.Context) {
	var account models.Account
	var response []models.DestinationResponse

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Println("Get destination failed: Failed converting id to integer", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Failed converting id to integer"})
		return
	}

	database.DATABASE.Find(&account, id)
	database.DATABASE.Model(&account).Association("TransferDestination").Find(&account.TransferDestination)

	for _, account := range account.TransferDestination {
		response = append(response, models.DestinationResponse{
			AccountNumber: account.AccountNumber,
			FirstName: account.FirstName,
			LastName: account.LastName,
			Username: account.Username,
		})
	}

	c.JSON(http.StatusOK, gin.H{"destinations": response})
}

func Transfer(c *gin.Context) {
	var source models.Account
	var destination models.Account
	var transfer models.Transfer

	// Bind arriving json into a transfer model
	err := c.BindJSON(&transfer)
	if err != nil {
		log.Println("Transfer failed: Failed binding json", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Failed binding request"})
		return
	}

	// Calculate conversion rate
	_, rate := getRatesFromRedis(transfer.Currency)
	conversion := float64(transfer.Amount) / rate
	newAmount := int64(math.Floor(conversion))
	transfer.ConvertedAmount = newAmount

	// Pull data from accounts table
	create := database.DATABASE.Create(&transfer)
	if create.Error != nil {
		log.Println("Transfer failed: Failed inserting to database", create.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Failed inserting transfer to database"})
		return
	}
	database.DATABASE.Find(&source, transfer.SourceID)
	database.DATABASE.Where("account_number=?", transfer.Destination).First(&destination)

	// If balance is insufficient
	if source.Balance < newAmount {
		database.DATABASE.Model(&transfer).Update("status", "failed")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
		return
	}

	// Subtract balance from source
	// Add balance to destination
	newSourceBalance := source.Balance - newAmount
	newDestinationBalance := destination.Balance + newAmount

	// Update database values
	database.DATABASE.Find(&source, transfer.SourceID).Update("balance", newSourceBalance)
	database.DATABASE.Find(&destination, transfer.Destination).Update("balance", newDestinationBalance)
	database.DATABASE.Model(&transfer).Update("status", "success")
	c.JSON(http.StatusOK, gin.H{"message": "Transfer completed"})
}