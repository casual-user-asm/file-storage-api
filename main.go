package main

import (
	"file-storage-api/config"
	"file-storage-api/storage"
	"fmt"
)

func init() {
	config.LoadEnv()
	storage.InitMinio()
}

func main() {
	err := storage.UploadFile("test", "/mnt/c/Users/Vladick/Downloads/test.jpg")
	if err != nil {
		fmt.Printf("Error upload the file: %v", err)
	}

	err = storage.DownloadFile("test", "test", "/mnt/c/Users/Vladick/Pictures/download.jpg")
	if err != nil {
		fmt.Printf("Error download the file: %v", err)
	}
}
