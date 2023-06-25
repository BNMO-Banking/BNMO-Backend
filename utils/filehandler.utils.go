package utils

import (
	"fmt"
	"mime/multipart"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var fileRoot = "./images"
var server_host = os.Getenv("SERVER_HOST")
var server_port = os.Getenv("SERVER_PORT")

func SaveFile(c *gin.Context, file *multipart.FileHeader) string {
	fileName := uuid.New().String() + file.Filename
	filePath := fmt.Sprintf("%s/id_cards/%s", fileRoot, fileName)
	err := c.SaveUploadedFile(file, filePath)
	if err != nil {
		HandleInternalServerError(c, err, "Save File", "Failed to save file")
	}

	return fmt.Sprintf("http://%s:%s/images/id_cards/%s", server_host, server_port, fileName)
}
