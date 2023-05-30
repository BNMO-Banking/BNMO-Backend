package controller

import (
	"BNMO/database"
	"BNMO/models"
	"log"
	"net/http"
	"os"
	"strconv"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ChangeImage(c *gin.Context) {
	var request models.ChangeImageRequest
	var account models.Account

	err := c.Bind(&request)
	if err != nil {
		log.Println("Change image failed: Failed binding form data", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Failed binding request"})
		return
	}

	id, err := strconv.Atoi(request.Id)
	if err != nil {
		log.Println("Change image failed: Failed converting id to integer", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Failed converting id to integer"})
		return
	}
	database.DATABASE.Find(&account, id)

	// Delete old file
	err = os.Remove("./images/" + request.OldName)
	if e, ok := err.(*os.PathError); ok && e.Err != syscall.ENOENT {
		log.Println("Change image failed: Failed in removing old file", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Failed in removing old file"})
		return
	}

	// Write image to folder
	fileName := uuid.New().String() + request.NewImage.Filename
	filePath := "./images/" + fileName
	err = c.SaveUploadedFile(request.NewImage, filePath)
	if err != nil {
		log.Println("Change image failed: Failed in saving file", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Failed in saving file"})
		return
	}

	storedFilePath := "http://localhost:8080/images/" + fileName
	account.ImagePath = storedFilePath

	database.DATABASE.Save(&account)
	c.JSON(http.StatusOK, gin.H{"message": "Image successfully changed", "image": storedFilePath})
}

func CheckPin(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var account models.Account

	database.DATABASE.Find(&account, id)
	if len(account.Pin) != 0 {
		c.JSON(http.StatusOK, gin.H{"pin_set": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"pin_set": false})
	}
}

func ChangePin(c * gin.Context) {
	var request models.ChangePinRequest
	var account models.Account

	err := c.BindJSON(&request)
	if err != nil {
		log.Println("Change PIN failed: Failed binding json", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Failed binding request"})
		return
	}

	database.DATABASE.Find(&account, request.Id)
	if len(account.Pin) != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "You can only change your PIN once. Contact our administrator to change it again"})
		return
	}

	if request.Pin != request.RepeatPin {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mismatched PIN"})
		return
	}

	account.SetPin(request.Pin)
	database.DATABASE.Save(&account)
	c.JSON(http.StatusOK, gin.H{"message": "PIN successfully set"})
}

func ChangePassword(c *gin.Context) {
	var request models.ChangePassRequest
	var account models.Account

	err := c.BindJSON(&request)
	if err != nil {
		log.Println("Change password failed: Failed binding json", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Failed binding request"})
		return
	}

	database.DATABASE.Find(&account, request.Id)
	err = account.ComparePassword(request.OldPass)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect old password"})
		return
	}

	if len(request.NewPass) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "New password too short"})
		return
	}

	if request.NewPass != request.RepeatPass {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mismatched password"})
		return
	}

	account.SetPassword(request.RepeatPass)
	database.DATABASE.Save(&account)
	c.JSON(http.StatusOK, gin.H{"message": "Password successfully changed"})
}

func SendAllCustomerData(c *gin.Context) {
	var customerDatas []models.Account
	var formattedDatas []map[string]interface{}

	// Pull data from the requests table inside the database
	database.DATABASE.Where(map[string]interface{}{"account_status": "accepted", "is_admin": false}).Find(&customerDatas)

	for _, data := range customerDatas {
		formattedDatas = append(formattedDatas, gin.H{
			"ID": data.ID,
			"CreatedAt": data.CreatedAt,
			"first_name": data.FirstName,
			"last_name": data.LastName,
			"email": data.Email,
			"username": data.Username,
			"image_path": data.ImagePath,
			"account_number": data.AccountNumber,
			"balance": data.Balance,
		})
	}

	// Return data to frontend
	c.JSON(http.StatusOK, gin.H{
		"data": formattedDatas,
	})
}

func EditData(c *gin.Context) {
	var request models.EditDataRequest
	var account models.Account

	err := c.BindJSON(&request)
	if err != nil {
		log.Println("Edit data failed: Failed binding json", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Failed binding request"})
		return
	}

	database.DATABASE.Find(&account, request.Id)

	// Update account data
	account.FirstName = request.FirstName
	account.LastName = request.LastName
	account.Email = request.Email
	account.Username = request.Username
	account.Balance = request.Balance

	database.DATABASE.Save(&account)
	c.JSON(http.StatusOK, gin.H{"message": "Data successfully edited"})
}

func ResetPIN(c *gin.Context) {
	var request models.ResetPinRequest
	var account models.Account
	err := c.BindJSON(&request)
	if err != nil {
		log.Println("Reset PIN failed: Failed binding json", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: Failed binding request"})
		return
	}

	database.DATABASE.Find(&account, request.Id)
	account.Pin = nil
	database.DATABASE.Save(&account)
	c.JSON(http.StatusOK, gin.H{"message": "PIN reset successful"})
}

func DeleteAccount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var account models.Account

	database.DATABASE.Find(&account, id)
	database.DATABASE.Delete(&account)
	c.JSON(http.StatusOK, gin.H{"message": "Delete account successful"})
}