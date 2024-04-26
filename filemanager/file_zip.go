package filemanager

import (
	"io"
	"os"
	"archive/zip"
	"path/filepath"
)

type FileZip struct {}

// ZipFileOrFolder compresses a file or folder into a ZIP archive.
// It takes a source file or folder path and a target ZIP file path as input.
// Returns an error if any occurs during the compression process.
func (fzip *FileZip) ZipFileOrFolder(source string, target string) error {
	// 1. Create a ZIP file and zip.Writer
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	// Go through all the files of the source
	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Create a local file header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		// set compression
		header.Method = zip.Deflate

		// Set relative path of a file as the header name
		header.Name, err = filepath.Rel(filepath.Dir(source), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		}
		// Create writer for the file header and save content of the file
		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(headerWriter, f)
		return err
	})
}
