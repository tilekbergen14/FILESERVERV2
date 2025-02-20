package routes

import (
	"fileserver/handlers"

	"github.com/gin-gonic/gin"
)

func Fileserver(router *gin.Engine) {
	router.POST("/sendfile", handlers.UploadFileHandler)
	router.POST("/download", handlers.DownloadFileHandler)
	router.POST("/clientdownload", handlers.DownloadFileHandler)

	router.GET("/serve", handlers.ServeFileHandler) 
}

