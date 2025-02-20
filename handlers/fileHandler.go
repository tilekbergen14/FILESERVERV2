package handlers

import (
	"fileserver/services"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UploadFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	fileContent, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read file"})
		return
	}
	defer fileContent.Close()

	customUuid := c.Request.PostFormValue("customUuid")
	subRootFolder := c.Request.PostFormValue("subrootfolder")
	filePath, err := services.UploadFile(subRootFolder, customUuid, fileContent, file.Filename)
	fmt.Println(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	c.JSON(http.StatusOK, filePath)
}

func DownloadFileHandler(c *gin.Context) {
	type DownloadRequest struct {
		FilePath string `json:"filePath"`
	}
	var request DownloadRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if request.FilePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File path is required"})
		return
	}

	file, err := services.GetFile("destination", request.FilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get file: %v", err)})
		return
	}

	fileContentType := http.DetectContentType(file)
	fileSize := len(file)

	c.Header("Content-Type", fileContentType)
	c.Header("Content-Length", fmt.Sprintf("%d", fileSize))
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", request.FilePath))
	c.Header("X-Content-Type-Options", "nosniff")

	_, err = c.Writer.Write(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to send file: %v", err)})
		return
	}
}

func ServeFileHandler(c *gin.Context) {
    filePath := c.DefaultQuery("path", "")
    if filePath == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Path parameter is required"})
        return
    }

    filePath = strings.Trim(filePath, "\"")

   	file, contentType, err := services.ServeFile("destination", filePath) 
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	c.Header("Content-Type", contentType)
	c.Status(http.StatusOK)
	c.Writer.Write(file) 
}
