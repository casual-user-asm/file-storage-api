package main

import (
	"file-storage-api/config"
	"file-storage-api/handlers"
	"file-storage-api/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	storage.InitMinio()
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("", func(c *gin.Context) {
		files, err := storage.ListAllFiles()
		if err != nil {
			c.String(http.StatusBadRequest, "Error: %v", err)
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main page for files",
			"files": files,
		})
	})
	router.POST("/upload", handlers.UploadFileHandler)
	router.GET("/files/:file", handlers.DownLoadFileHandler)
	router.Run(":8080")
}
