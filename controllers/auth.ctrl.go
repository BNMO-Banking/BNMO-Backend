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
		utils.HandleRecordNotFound(c, "Register", "Email already exist")
		return
	}

	// Validate phone number
	if !utils.ValidatePhoneNumber(request.PhoneNumber) {
		utils.HandleBadRequest(c, "Register", "Invalid phone number")
		return
	}

	// Validate username availability
	err = database.DB.Where("username = ?", request.Username).First(&account).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.HandleRecordNotFound(c, "Register", "Username already taken")
		return
	}

	// File handling
	filePath := utils.SaveFile(c, request.IdCard, enum.FILE_ID_CARD)
	if len(filePath) == 0 {
		return
	}

	// Hashing password
	password, err := utils.HashPassword(request.Password)
	if err != nil {
		utils.HandleInternalServerError(c, err, "Register", "Failed to hash password")
		return
	}

	// Create new data entry
	newAccount := gormmodels.Customer{
		PhoneNumber: request.PhoneNumber,
		IdCardPath:  filePath,
		Address: gormmodels.CustomerAddress{
			AddressLine1: request.AddressLine1,
			AddressLine2: request.AddressLine2,
			City:         request.City,
			State:        request.State,
			PostalCode:   request.PostalCode,
			Country:      request.Country,
		},
		Account: gormmodels.Account{
			FirstName:   request.FirstName,
			LastName:    request.LastName,
			Email:       request.Email,
			Username:    request.Username,
			Password:    password,
			AccountType: enum.CUSTOMER,
		},
	}

	database.DB.Create(&newAccount)
	c.JSON(http.StatusOK, gin.H{"message": "Registration successful. Please wait for validation"})
}

func LoginAccount(c *gin.Context) {
	var request models.LoginReq
	var account gormmodels.Account

	err := c.BindJSON(&request)
	if err != nil {
		utils.HandleInternalServerError(c, err, "Login", "Failed to bind request")
		return
	}

	// Fetch account
	err = database.DB.Where("email = ?", request.EmailUsername).Or("username = ?", request.EmailUsername).First(&account).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.HandleRecordNotFound(c, "Login", "Email / username is incorrect")
		return
	}

	// Compare password
	err = utils.ComparePassword(account.Password, request.Password)
	if err != nil {
		utils.HandleBadRequest(c, "Login", "Incorrect password")
		return
	}

	// Check if admin or customer
	if account.AccountType == enum.ADMIN {
		var admin gormmodels.Admin

		database.DB.Preload("Account").Where("account_id = ?", account.ID).First(&admin)

		token, err := token.GenerateJWT(account.ID.String(), account.AccountType)
		if err != nil {
			utils.HandleInternalServerError(c, err, "Login", "Failed to generate token")
			return
		}

		response := models.LoginResAccount{
			Email:       account.Email,
			Username:    account.Username,
			AccountType: account.AccountType,
			AccountRole: admin.Role,
		}

		c.JSON(http.StatusOK, models.LoginRes{
			Account: response,
			Token:   token,
		},
		)
	} else if account.AccountType == enum.CUSTOMER {
		var customer gormmodels.Customer
		// Check account validation status
		database.DB.Preload("Account").Where("account_id = ?", account.ID).First(&customer)
		if customer.Status == enum.ACCOUNT_ACCEPTED {
			token, err := token.GenerateJWT(account.ID.String(), account.AccountType)
			if err != nil {
				utils.HandleInternalServerError(c, err, "Login", "Failed to generate token")
				return
			}

			response := models.LoginResAccount{
				Id:          customer.ID,
				Email:       account.Email,
				Username:    account.Username,
				AccountType: account.AccountType,
			}

			var pinStatus enum.PinStatus
			if len(customer.AccountNumber) == 0 {
				pinStatus = enum.PIN_UNSET
			} else {
				pinStatus = enum.PIN_SET
			}

			c.JSON(http.StatusOK, models.LoginRes{
				Account:   response,
				PinStatus: pinStatus,
				Token:     token,
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
