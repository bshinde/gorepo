package controller

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/scanner/internal/downloader"
	"github.com/scanner/pkg/utils"
)

// ProcessGitRepo handles the Git repository processing and clones it.
func ProcessGitRepo(url string) error {
	// Define the local cache directory
	cacheDir := getCacheDirectory()

	// Determine the directory name to store the repo
	dirName := utils.GetDirectoryNameFromURL(url)
	dirPath := filepath.Join(cacheDir, dirName)

	// Clone the repository into the cache directory
	err := downloader.CloneRepo(url, dirPath)
	if err != nil {
		log.Printf("Failed to clone %s: %v\n", url, err)
		return err
	}

	fmt.Printf("Cloned Git repository to: %s\n", dirPath)
	return nil
}

// ProcessFile handles the file system processing (downloading the file).
func ProcessFile(url string) error {
	// Define the local cache directory
	cacheDir := getCacheDirectory()

	// Determine the directory name to store the file
	dirName := utils.GetFileNameFromURL(url)
	dirPath := filepath.Join(cacheDir, dirName)

	// Download the file into the cache directory
	err := downloader.DownloadFile(url, dirPath)
	if err != nil {
		log.Printf("Failed to download file %s: %v\n", url, err)
		return err
	}

	fmt.Printf("Downloaded file to: %s\n", dirPath)
	return nil
}

// getCacheDirectory returns the local cache directory path
func getCacheDirectory() string {
	cacheDir := "./cache" // Default cache directory
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		err := os.MkdirAll(cacheDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Error creating cache directory: %v", err)
		}
	}
	return cacheDir
}
