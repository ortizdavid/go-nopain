package main

import (
  "fmt"
  "net/http"
	"time"
  "github.com/ortizdavid/go-nopain/filemanager" 
)

func main() {
	// Define upload configurations
	destinationPath := "./uploads"
	maxSize  := 5  // Maximum file size in MB
	allowedExtensions := filemanager.ExtImages // Allow only image files

	// Create a new Uploader instance
	uploader := filemanager.NewUploader(destinationPath, maxSize, allowedExtensions)

	// Define an HTTP handler function for file upload
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
		fmt.Fprintf(w, "Only POST requests allowed")
		return
	}

    // Upload single file using UploadSingleFile
    fileInfo, err := uploader.UploadSingleFile(r, "file") // Assuming your form field name is "file"
    if err != nil {
      fmt.Fprintf(w, "Error uploading file: %v", err)
      return
    }

    fmt.Fprintf(w, "File uploaded successfully!\nDetails:\n  - Original name: %s\n  - Final name: %s\n  - Size: %d bytes\n  - Content type: %s\n  - Uploaded at: %s\n",
      fileInfo.OriginalFileName, fileInfo.FinalName, fileInfo.Size, fileInfo.ContentType, fileInfo.UploadTime.Format(time.RFC3339))
  })

  // Start the HTTP server
  fmt.Println("Server listening on port 8080")
  http.ListenAndServe(":8080", nil)
}

