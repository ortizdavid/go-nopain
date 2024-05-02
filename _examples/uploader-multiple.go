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

    // Upload multiple files using UploadMultipleFiles
    uploadInfos, err := uploader.UploadMultipleFiles(r, "files") // Assuming your form field name is "files" (for multiple files)
    if err != nil {
      fmt.Fprintf(w, "Error uploading files: %v", err)
      return
    }

    fmt.Fprintf(w, "Files uploaded successfully!\n")
    for _, info := range uploadInfos {
      fmt.Fprintf(w, "  - Original name: %s\n  - Final name: %s\n  - Size: %d bytes\n  - Content type: %s\n  - Uploaded at: %s\n",
        info.OriginalFileName, info.FinalName, info.Size, info.ContentType, info.UploadTime.Format(time.RFC3339))
    }
  })

  // Start the HTTP server
  fmt.Println("Server listening on port 8090")
  http.ListenAndServe(":8090", nil)
}
