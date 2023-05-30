package controller

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func generateRandomName(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Upload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println("Error in multipart form", err)
		return
	}

	files := form.File["image"]
	fileName := ""

	for _, file := range files {
		fileName = generateRandomName(5) + "-" + file.Filename
		err := c.SaveUploadedFile(file, "images/"+fileName)
		if err != nil {
			fmt.Println("Error in saving uploaded file", err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"url": "http://localhost:8080/api/uploads/" + fileName,
	})
}