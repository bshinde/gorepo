package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// IsGitRepository checks if the URL is a Git repository (ends with .git).
func IsGitRepository(url string) bool {
	return strings.HasSuffix(url, ".git")
}

// GetDirectoryNameFromURL generates a directory name based on the URL, without `.git` for repositories.
func GetDirectoryNameFromURL(url string) string {
	// If it's a Git repository, remove the ".git" suffix
	if IsGitRepository(url) {
		url = strings.TrimSuffix(url, ".git")
	}
	// Extract the directory name (e.g., repo1 for https://github.com/username/repo1.git)
	dirName := url[strings.LastIndex(url, "/")+1:]
	return dirName
}

// GetFileNameFromURL extracts the file name from the URL.
func GetFileNameFromURL(url string) string {
	// If the URL ends with a '/', strip it off
	if strings.HasSuffix(url, "/") {
		url = url[:len(url)-1]
	}
	// Extract the file name by taking the part after the last '/'
	return url[strings.LastIndex(url, "/")+1:]
}

// CreateTempDirectory creates a temporary directory to store Git repos or file system updates.
// func CreateTempDirectory() (string, error) {
// 	// Use a system temp directory for cache storage
// 	tempDir := os.TempDir()

// 	// Ensure the temp directory exists
// 	err := os.MkdirAll(tempDir, os.ModePerm)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to create temp directory: %w", err)
// 	}

//		return tempDir, nil
//	}
func CreateTempDirectory() (string, error) {
	// Get the system's default temporary directory
	tempDir := os.TempDir()

	// Define a hidden directory path (for Unix-based, prepend a dot to the folder name)
	// For Windows, we'll set the hidden and system attributes later
	var hiddenDir string
	if runtime.GOOS == "windows" {
		// On Windows, create a directory in Temp and mark it as hidden + system
		hiddenDir = filepath.Join(tempDir, "tempDri")
	} else {
		// On Unix-based systems, create a hidden directory by prefixing with a dot
		hiddenDir = filepath.Join(tempDir, ".tempDri")
	}

	// Create the directory (if it doesn't exist)
	err := os.MkdirAll(hiddenDir, 0700) // Use restrictive permissions (only accessible by the owner)
	if err != nil {
		return "", fmt.Errorf("failed to create hidden directory: %w", err)
	}

	// For Windows, set the "hidden" and "system" attributes using the attrib command
	if runtime.GOOS == "windows" {
		err := exec.Command("attrib", "+h", "+s", hiddenDir).Run() // +s marks it as a system folder
		if err != nil {
			return "", fmt.Errorf("failed to set hidden and system attributes on Windows: %w", err)
		}
	}

	return hiddenDir, nil
}

// GetTempRepoPath generates a path inside the temp directory for a specific repository or file.
func GetTempRepoPath(url string) (string, error) {
	// Get the temp directory
	tempDir, err := CreateTempDirectory()
	if err != nil {
		return "", err
	}

	// Get the directory name based on the URL (repository name)
	dirName := GetDirectoryNameFromURL(url)

	// Construct the full path where the repository will be stored
	repoPath := filepath.Join(tempDir, dirName)

	return repoPath, nil
}

// DownloadAndStoreGitRepo handles the downloading and storing of a Git repository to the temp directory.
func DownloadAndStoreGitRepo(url string) (string, error) {
	// Get the path where the repo will be stored in temp
	repoPath, err := GetTempRepoPath(url)
	if err != nil {
		return "", err
	}

	// Simulating the downloading of the repo (this can be replaced by actual git clone logic)
	// For now, let's create an empty file to indicate the repository is "stored"
	err = os.WriteFile(filepath.Join(repoPath, "README.md"), []byte("This is a placeholder for the Git repo."), 0644)
	if err != nil {
		return "", fmt.Errorf("failed to store the repository: %w", err)
	}

	return repoPath, nil
}
