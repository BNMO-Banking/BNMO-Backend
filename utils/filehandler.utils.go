package utils

import (
	"BNMO/enum"
	"fmt"
	"mime/multipart"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SaveFile(c *gin.Context, file *multipart.FileHeader, fileType enum.FileType) string {
	var fileRoot = "./images"
	var server_addr = os.Getenv("SERVER_ADDR")
	fileName := uuid.New().String() + file.Filename
	if fileType == enum.FILE_ID_CARD {
		filePath := fmt.Sprintf("%s/id_cards/%s", fileRoot, fileName)
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			HandleInternalServerError(c, err, "Save file", "Failed to save file")
			return ""
		}

		return fmt.Sprintf("%simages/id_cards/%s", server_addr, fileName)
	} else if fileType == enum.FILE_PROFILE_PICTURE {
		fileName := uuid.New().String() + file.Filename
		filePath := fmt.Sprintf("%s/profile_pics/%s", fileRoot, fileName)
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			HandleInternalServerError(c, err, "Save file", "Failed to save file")
			return ""
		}

		return fmt.Sprintf("%simages/profile_pics/%s", server_addr, fileName)
	} else {
		return ""
	}
}

func DeleteFile(filePath string) error {
	var server_addr = os.Getenv("SERVER_ADDR")
	fileSlice := strings.SplitN(filePath, server_addr, -1)
	fileName := fileSlice[1]
	err := os.Remove(fmt.Sprintf("./%s", fileName))
	if err != nil {
		return err
	}

	return nil
}
