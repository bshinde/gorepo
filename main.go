package main

import (
	"log"

	"github.com/scanner/pkg/uploader"
)

func main() {
	// Handle user input for repo/file upload
	err := uploader.UploadRequest()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
