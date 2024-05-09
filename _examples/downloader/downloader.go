package main

import (
	"fmt"
	"log"

	"github.com/ortizdavid/go-nopain/filemanager" 
)

func main() {
	// Create a new instance of Downloader
	downloader := filemanager.NewDownloader()

	// URL of the file to download
	url := "https://filesamples.com/samples/document/txt/sample3.txt" // Replace with your desired URL

	// Download the file
	downloadInfo, err := downloader.DownloadFile(url)
	if err != nil {
		log.Fatalf("Error downloading file: %v", err)
	}

	// Print download information
	fmt.Println("Download Information:")
	fmt.Printf("URL: %s\n", downloadInfo.URL)
	fmt.Printf("File Name: %s\n", downloadInfo.FileName)
	fmt.Printf("Directory: %s\n", downloadInfo.Directory)
	fmt.Printf("Size: %d MB\n",downloadInfo.Size)
}
