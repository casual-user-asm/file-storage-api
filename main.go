package main

import (
	"file-storage-api/config"
)

func init() {
	config.LoadEnv()
}
