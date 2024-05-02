package filemanager

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Downloader manages file downloads.
type Downloader struct {
	Errors []string
}

// DownloadInfo contains information about a downloaded file
type DownloadInfo struct {
	URL       string // URL is the source URL from which the file is downloaded.
	FileName  string // FileName is the name of the downloaded file.
	Directory string // Directory is the directory where the file is saved.
	Size	  int64	// Size is the size of the downloaded file in MB.
}

// NewDownloader creates a new instance of Downloader.
func NewDownloader() Downloader {
	return Downloader{
		Errors: make([]string, 0),
	}
}

// DownloadFile downloads a file from the specified URL.
// DownloadFile downloads a file from the specified URL.
func (dl Downloader) DownloadFile(url string) (DownloadInfo, error) {
	if !dl.isValidURL(url) {
		return DownloadInfo{}, fmt.Errorf("invalid URL: %s", url)
	}

	fileName := dl.extractNameFromURL(url)
	directory := "." // Default directory is current working directory
	fullPath := filepath.Join(directory, fileName)

	// Create directory if it doesn't exist
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return DownloadInfo{}, fmt.Errorf("error creating directory: %w", err)
	}

	// Download the file
	resp, err := http.Get(url)
	if err != nil {
		return DownloadInfo{}, fmt.Errorf("error downloading file: %w", err)
	}
	defer resp.Body.Close()

	// Check for successful download
	if resp.StatusCode != http.StatusOK {
		return DownloadInfo{}, fmt.Errorf("error downloading file: %s", resp.Status)
	}

	// Create the destination file
	out, err := os.Create(fullPath)
	if err != nil {
		return DownloadInfo{}, fmt.Errorf("error creating file: %w", err)
	}
	defer out.Close()

	// Write downloaded data to the file
	downloaded, err := io.Copy(out, resp.Body)
	if err != nil {
		return DownloadInfo{}, fmt.Errorf("error writing file: %w", err)
	}

	// Calculate and set file size in MB
	size := float64(downloaded) / (1 << 20) // Conversion to megabytes (MB)

	return DownloadInfo{
		URL:       url,
		FileName:  fileName,
		Directory: directory,
		Size:      int64(size), // Convert to int64
	}, nil
}

func (dl Downloader) isValidURL(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}

func (dl Downloader) extractNameFromURL(url string) string {
	return filepath.Base(url)
}