package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/scanner/pkg/utils"
)

// CloneRepo clones a Git repository from the given URL into a specified directory.
func CloneRepo(url, dirName string) error {
	// Create the directory if it doesn't exist
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	// Clone the Git repository into the directory
	cmd := exec.Command("git", "clone", url, dirName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// DownloadFile downloads a file from the given URL into a specified directory.
func DownloadFile(url, dirName string) error {
	// Create the directory if it doesn't exist
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	// Get the data from the URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file in the specified directory
	fileName := fmt.Sprintf("%s/%s", dirName, utils.GetFileNameFromURL(url))
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Copy the response body into the file
	_, err = io.Copy(file, resp.Body)
	return err
}
