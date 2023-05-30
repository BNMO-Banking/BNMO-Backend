package controller

import (
	"BNMO/database"
	"BNMO/models"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func UpdateImage(c *gin.Context) {
	var data map[string]interface{}
	var account models.Account

	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("Unable to parse body into a validate_request struct:" + err.Error())
		return
	}

	// Delete the old image
	oldUrl := data["old_url"].(string)
	oldFormat := oldUrl[34:]
	err = os.Remove("./images/" + oldFormat)
	if err != nil {
		fmt.Println("Failed to delete file" + err.Error())
		return
	}
	
	// Access account tables on database based on the id and change the url
	database.DATABASE.First(&account, uint(data["id"].(float64))).Update("image_path", data["new_url"].(string))
	c.JSON(http.StatusOK, gin.H{"message": "Image updated successfully"})
}