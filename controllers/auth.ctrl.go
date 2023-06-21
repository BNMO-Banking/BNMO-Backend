package controllers

import (
	"BNMO/database"
	"BNMO/models"
	"BNMO/token"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9. %+\-]+@[a-z0-9. %+\-]+\.[a-z0-9. %+\-]`)
	return Re.MatchString(email)
}

func generateAccountNumber() string {
	firstSequel := strconv.Itoa(rand.Intn(1000))
	if len(firstSequel) < 3 {
		firstSequel = "0" + firstSequel
	} else if len(firstSequel) < 2 {
		firstSequel = "00" + firstSequel
	}

	secondSequel := strconv.Itoa(rand.Intn(1000))
	if len(secondSequel) < 3 {
		secondSequel = "0" + secondSequel
	} else if len(secondSequel) < 2 {
		secondSequel = "00" + secondSequel
	}

	thirdSequel := strconv.Itoa(rand.Intn(10000))
	if len(thirdSequel) < 4 {
		thirdSequel = "0" + thirdSequel
	} else if len(thirdSequel) < 3 {
		thirdSequel = "00" + thirdSequel
	} else if len(thirdSequel) < 2 {
		thirdSequel = "000" + thirdSequel
	}

	return fmt.Sprintf("%s-%s-%s", firstSequel, secondSequel, thirdSequel)
}

func RegisterAccount(c *gin.Context) {
	var request models.RegisterRequest
	var accountData models.Account

	// Bind arriving json into register model
	err := c.Bind(&request)
	if err != nil {
		log.Println("Register failed: Failed binding form data", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Failed binding request"})
		return
	}

	// Check if the length of password is less than 8 characters
	if len(request.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be 8 characters or more"})
		return
	}

	// Validate the email
	if !validateEmail(request.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address"})
		return
	}

	// Check if email already exist within the database
	database.DATABASE.Where("email=?", request.Email).First(&accountData)
	if accountData.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	// Check if username already exist within the database
	database.DATABASE.Where("username=?", request.Username).First(&accountData)
	if accountData.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	// Write image to folder
	fileName := uuid.New().String() + request.Image.Filename
	filePath := "./images/" + fileName
	err = c.SaveUploadedFile(request.Image, filePath)
	if err != nil {
		log.Println("Register failed: Failed in saving file", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Failed in saving file"})
		return
	}

	storedFilePath := "http://localhost:8080/images/" + fileName

	// Insert register data to account model
	account := models.Account{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Username:  request.Username,
		ImagePath: storedFilePath,
	}

	// Hash password using bcrypt
	account.SetPassword(request.Password)

	// Insert the data into the database
	insert := database.DATABASE.Create(&account)
	if insert.Error != nil {
		log.Println("Register failed: Failed inserting to database", insert.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Failed to insert account to database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Account successfully registered. Please wait for validation."})
}

func LoginAccount(c *gin.Context) {
	var request models.LoginRequest
	var account models.Account

	// Bind arriving json into login model
	err := c.BindJSON(&request)
	if err != nil {
		log.Println("Login failed: Failed binding json", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Failed binding request"})
		return
	}

	// Check if email exists inside the database
	database.DATABASE.Where("email=?", request.Email).First(&account)
	if account.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email does not exist"})
		return
	}

	// Check password validity
	err = account.ComparePassword(request.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password"})
		return
	}

	if account.AccountStatus.String == "accepted" {
		// Authenticate user
		if err != nil {
			log.Println("Login failed: Failed generating JWT", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Failed to generate JWT"})
			return
		}

		token, err := token.GenerateJWT(strconv.Itoa(int(account.ID)))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Error generating token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"account": gin.H{
				"ID":             account.ID,
				"is_admin":       account.IsAdmin.Bool,
				"first_name":     account.FirstName,
				"last_name":      account.LastName,
				"email":          account.Email,
				"username":       account.Username,
				"image_path":     account.ImagePath,
				"account_number": account.AccountNumber,
				"balance":        account.Balance,
				"CreatedAt":      account.CreatedAt,
			},
			"token":         token,
			"accountStatus": account.AccountStatus.String,
			"message":       "Login successful"})
	} else if account.AccountStatus.String == "pending" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Account isn't verified. Please wait for validation"})
		return
	} else if account.AccountStatus.String == "rejected" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Account is rejected. Please contact our support"})
		return
	}
}

func LogoutAccount(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Log out successful"})
}