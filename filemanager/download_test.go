package filemanager

import "testing"


var downloader = NewDownloader()

func Test_Download_NonExistingURL(t *testing.T) {
	nonExistingURL := "http://localhost:12345/nonexistentfile.txt"
	
	_, err := downloader.DownloadFile(nonExistingURL)
	if err == nil {
		t.Errorf("This Download must Fail. ")
	}
}


func Test_Download_ExistingUrl(t *testing.T) {
	existingURL := "https://filesamples.com/samples/document/txt/sample3.txt"
	
	_, err := downloader.DownloadFile(existingURL)
	if err != nil {
		t.Errorf("Error for existing URL: %s. ", existingURL)
	}
}