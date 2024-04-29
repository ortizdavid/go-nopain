package filemanager

type Downloader struct {
	Errors []string
}

func NewDownloader() Downloader {
	return Downloader{}
}

func (dl Downloader) DownloadFile() error {
	return nil
}