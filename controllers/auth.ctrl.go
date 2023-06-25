package controllers

import (
	"BNMO/database"
	"BNMO/enum"
	gormmodels "BNMO/gorm_models"
	"BNMO/models"
	"BNMO/utils"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	salt, err := strconv.Atoi(os.Getenv("PASS_SALT"))
	if err != nil {
		utils.HandleInternalServerError(c, err, "Register", "Failed to get salt")
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), salt)
	if err != nil {
		utils.HandleInternalServerError(c, err, "Register", "Failed to hash password")
		return
	}

	// Create new data entry
	newAccount := gormmodels.Customer{
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
	}

	database.DB.Create(newAccount)
	c.JSON(http.StatusOK, gin.H{"message": "Registration successful. Please wait for validation"})
}

// func LoginAccount(c *gin.Context) {
// 	var request models.LoginRequest
// 	var account models.Account

// 	// Bind arriving json into login model
// 	err := c.BindJSON(&request)
// 	if err != nil {
// 		log.Println("Login failed: Failed binding json", err.Error())
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Failed binding request"})
// 		return
// 	}

// 	// Check if email exists inside the database
// 	database.DATABASE.Where("email=?", request.Email).First(&account)
// 	if account.ID == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Email does not exist"})
// 		return
// 	}

// 	// Check password validity
// 	err = account.ComparePassword(request.Password)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password"})
// 		return
// 	}

// 	if account.AccountStatus.String == "accepted" {
// 		// Authenticate user
// 		if err != nil {
// 			log.Println("Login failed: Failed generating JWT", err.Error())
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Failed to generate JWT"})
// 			return
// 		}

// 		token, err := token.GenerateJWT(strconv.Itoa(int(account.ID)))
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Error generating token"})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{
// 			"account": gin.H{
// 				"ID":             account.ID,
// 				"is_admin":       account.IsAdmin.Bool,
// 				"first_name":     account.FirstName,
// 				"last_name":      account.LastName,
// 				"email":          account.Email,
// 				"username":       account.Username,
// 				"image_path":     account.ImagePath,
// 				"account_number": account.AccountNumber,
// 				"balance":        account.Balance,
// 				"CreatedAt":      account.CreatedAt,
// 			},
// 			"token":         token,
// 			"accountStatus": account.AccountStatus.String,
// 			"message":       "Login successful"})
// 	} else if account.AccountStatus.String == "pending" {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Account isn't verified. Please wait for validation"})
// 		return
// 	} else if account.AccountStatus.String == "rejected" {
// 		c.JSON(http.StatusForbidden, gin.H{"error": "Account is rejected. Please contact our support"})
// 		return
// 	}
// }

func LoginAccount(c *gin.Context)

func LogoutAccount(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Log out successful"})
}
