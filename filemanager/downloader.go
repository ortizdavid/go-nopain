package filemanager

// Downloader manages file downloads.
type Downloader struct {
	Errors []string
}

// DownloadInfo contains information about a downloaded file
type DownloadInfo struct {
	URL       string // URL is the source URL from which the file is downloaded.
	FileName  string // FileName is the name of the downloaded file.
	Directory string // Directory is the directory where the file is saved.
}

// NewDownloader creates a new instance of Downloader.
func NewDownloader() Downloader {
	return Downloader{
		Errors: make([]string, 0),
	}
}

// DownloadFile downloads a file from the specified URL.
func (dl Downloader) DownloadFile(url string) (DownloadInfo, error) {
	return DownloadInfo{}, nil
}


func (dl Downloader) isValidURL(url string) bool {
	return false
}

func (dl Downloader) exctarctNameFromURL(url string) string {
	return ""
}