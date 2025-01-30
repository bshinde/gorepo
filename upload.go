package uploader

import (
	"fmt"

	"github.com/scanner/controller"
	"github.com/scanner/pkg/utils"
)

// UploadRequest handles the user request to upload a Git repository or file.
func UploadRequest() error {
	var input string
	fmt.Print("Enter the URL (Git repository or file): ")
	_, err := fmt.Scanln(&input)
	if err != nil || input == "" {
		return fmt.Errorf("invalid input: %v", err)
	}

	// Check if it's a Git repository or file system
	if utils.IsGitRepository(input) {
		// It's a Git repository
		fmt.Println("Detected Git repository. Proceeding to clone...")
		return controller.ProcessGitRepo(input)
	} else {
		// It's a file system (URL or local path)
		fmt.Println("Detected file system. Proceeding to download...")
		return controller.ProcessFile(input)
	}
}
