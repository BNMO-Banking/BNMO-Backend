package controllers

import (
	"BNMO/database"
	"BNMO/models"
	"database/sql"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)



func UpdateBalance(c *gin.Context) {
	var account models.Account
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Println("Get destination failed: Failed converting id to integer", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Failed converting id to integer"})
		return
	}
	
	database.DATABASE.Find(&account, id)
	c.JSON(http.StatusOK, gin.H{
		"balance": account.Balance,
	})
}

type Claims struct {
	jwt.StandardClaims
} 

func DisplayPendingAccount(c *gin.Context) {
	// Specify limitations
	page, _ := strconv.Atoi(c.Query("page"))
	limit := 5
	offset := (page-1) * limit

	var total int64
	var getAccounts []models.Account

	// Pull data from the requests table inside the database
	// Pull only based on the number of offsets and limits specified
	database.DATABASE.Offset(offset).Limit(limit).Where("account_status=?", "pending").Find(&getAccounts)
	database.DATABASE.Model(&models.Account{}).Where("account_status=?", "pending").Count(&total)

	// Return data to frontend
	c.JSON(http.StatusOK, gin.H{
		"data": getAccounts,
		"metadata": gin.H{
			"total": total,
			"page": page,
			"last_page": math.Ceil(float64(total)/float64(limit)),
		},
	})
}

func ValidateAccount(c *gin.Context) {
	var request models.ValidateAccount
	var account models.Account

	// Bind arriving json into validate account model
	err := c.BindJSON(&request)
	if err != nil {
		log.Println("Validate account failed: Failed binding json", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Failed binding request"})
		return
	}

	if request.Validation == "accepted" {
		database.DATABASE.First(&account, request.Id).Updates(models.Account{
			AccountStatus: sql.NullString{String: "accepted", Valid: true},
			AccountNumber: generateAccountNumber(),
		})
		c.JSON(http.StatusOK, gin.H{"message": "Account successfully accepted"})
		return
	} else if request.Validation == "rejected" {
		database.DATABASE.First(&account, request.Id).Update("account_status", "rejected")
		c.JSON(http.StatusOK, gin.H{"message": "Account successfully rejected"})
		return
	}

	
}
