package handlers

import (
	"file-storage-api/storage"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFileHandler(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "Error: %v", err)
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.String(http.StatusBadRequest, "Err: %v", err)
		return
	}
	defer file.Close()

	buf := make([]byte, 512)
	n, _ := file.Read(buf)
	contentType := http.DetectContentType(buf[:n])

	file.Seek(0, io.SeekStart)

	objectName := fileHeader.Filename
	err = storage.UploadFile(objectName, file, fileHeader.Size, contentType)
	if err != nil {
		c.String(http.StatusBadRequest, "Error: %v", err)
		return
	}

	c.HTML(http.StatusOK, "upload.html", gin.H{
		"message": fmt.Sprintf("Uploaded %s successfully", objectName),
	})
}

func DownLoadFileHandler(c *gin.Context) {
	fileName := c.Param("file")

	err := storage.DownloadFile(fileName)
	if err != nil {
		c.String(http.StatusBadRequest, "Error: %v", err)
		return
	}

	c.String(http.StatusOK, "File downloaded successfully!")
}
