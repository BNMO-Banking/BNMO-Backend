package controllers

import (
	"BNMO/database"
	"BNMO/enum"
	gormmodels "BNMO/gorm_models"
	"BNMO/models"
	"BNMO/token"
	"BNMO/utils"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAccount(c *gin.Context) {
	var request models.RegisterReq
	var account gormmodels.Account

	err := c.Bind(&request)
	if err != nil {
		utils.HandleInternalServerError(c, err, "Register", "Failed to bind request")
		return
	}

	// Validate password
	if len(request.Password) < 8 {
		utils.HandleBadRequest(c, "Register", "Password too short")
		return
	} else if strings.Compare(request.Password, request.ConfirmPassword) != 0 {
		utils.HandleBadRequest(c, "Register", "Confirm password do not match")
		return
	}

	// Validate email and availability
	if !utils.ValidateEmail(request.Email) {
		utils.HandleBadRequest(c, "Register", "Invalid email")
		return
	}

	err = database.DB.Where("email=?", request.Email).First(&account).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.HandleBadRequest(c, "Register", "Email already exist")
		return
	}

	// Validate phone number
	if !utils.ValidatePhoneNumber(request.PhoneNumber) {
		utils.HandleBadRequest(c, "Register", "Invalid phone number")
		return
	}

	// Validate username availability
	err = database.DB.Where("username=?", request.Username).First(&account).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.HandleBadRequest(c, "Register", "Username already taken")
		return
	}

	// File handling
	filePath := utils.SaveFile(c, request.IdCard)

	// Hashing password
	password, err := utils.HashPassword(request.Password)
	if err != nil {
		utils.HandleInternalServerError(c, err, "Register", "Failed to hash password")
	}

	// Create new data entry
	newAccount := gormmodels.CustomerAddress{
		AddressLine1: request.AddressLine1,
		AddressLine2: request.AddressLine2,
		State:        request.State,
		PostalCode:   request.PostalCode,
		Country:      request.Country,
		Customer: gormmodels.Customer{
			Status:      enum.ACCOUNT_PENDING,
			PhoneNumber: request.PhoneNumber,
			IdCardPath:  filePath,
			Account: gormmodels.Account{
				FirstName:   request.FirstName,
				LastName:    request.LastName,
				Email:       request.Email,
				Username:    request.Username,
				Password:    password,
				AccountType: enum.CUSTOMER,
			},
		},
	}

	database.DB.Create(newAccount)
	c.JSON(http.StatusOK, gin.H{"message": "Registration successful. Please wait for validation"})
}

func LoginAccount(c *gin.Context) {
	var request models.LoginReq
	var account gormmodels.Account

	// Bind arriving json into login model
	err := c.BindJSON(&request)
	if err != nil {
		utils.HandleBadRequest(c, "Login", "Failed to bind request")
		return
	}

	// Fetch account
	err = database.DB.Where("email=? OR username=?", request.EmailUsername).First(&account).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.HandleBadRequest(c, "Login", "Email / username is incorrect")
		return
	}

	// Compare password
	err = utils.ComparePassword(request.Password, account.Password)
	if err != nil {
		utils.HandleBadRequest(c, "Login", "Incorrect password")
	}

	// Check if admin or customer
	if account.AccountType == enum.ADMIN {
		var admin gormmodels.Admin

		database.DB.Where("account_id?=", account.ID).First(&admin)

		token, err := token.GenerateJWT(account.ID.String())
		if err != nil {
			utils.HandleInternalServerError(c, err, "Login", "Failed to generate token")
		}

		resAccount := models.LoginResAccount{
			Email:       account.Email,
			Username:    account.Username,
			AccountType: account.AccountType,
			AccountRole: admin.Role,
		}
		c.JSON(http.StatusOK, models.LoginRes{
			Account: resAccount,
			Token:   token,
		},
		)
	} else if account.AccountType == enum.CUSTOMER {
		var customer gormmodels.Customer
		// Check account validation status
		database.DB.Where("account_id=?", account.ID).First(&customer)
		if customer.Status == enum.ACCOUNT_ACCEPTED {
			token, err := token.GenerateJWT(account.ID.String())
			if err != nil {
				utils.HandleInternalServerError(c, err, "Login", "Failed to generate token")
			}

			resAccount := models.LoginResAccount{
				Email:       account.Email,
				Username:    account.Username,
				AccountType: account.AccountType,
			}
			c.JSON(http.StatusOK, models.LoginRes{
				Account: resAccount,
				Token:   token,
			},
			)
		} else if customer.Status == enum.ACCOUNT_PENDING {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Account isn't verified. Please wait for validation"})
			return
		} else if customer.Status == enum.ACCOUNT_REJECTED {
			c.JSON(http.StatusForbidden, gin.H{"error": "Account is rejected. Please contact our support"})
			return
		}
	}
}

func LogoutAccount(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Log out successful"})
}
