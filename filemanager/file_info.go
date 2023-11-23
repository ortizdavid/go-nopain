package filemanager

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
	"path/filepath"
	"github.com/ortizdavid/go-nopain/collections"
)

type FileInfo struct {}


func (finfo *FileInfo) GetFileInfo(fileName string) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("\nFile Name:", fileInfo.Name())  
	fmt.Println("Extension:", finfo.GetFileExt(fileName))          
	fmt.Println("Size:", fileInfo.Size(), " bytes")  
	fmt.Println("Type:", finfo.GetFileType(finfo.GetFileExt(fileName)))              
	fmt.Println("Last Modified:", fileInfo.ModTime()) 
	fmt.Println("Permissions:", fileInfo.Mode())     
}


func (finfo *FileInfo) ListFiles(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("\nAll Files in '%s':\n", dir)
	printChar("-", 125)
	fmt.Println("NAME\t\t\t\tTYPE\t\t\tSIZE(Bytes)\t\t\tLAST MODIFIED")
	printChar("-", 125)
	for _, file := range files {
		ext := finfo.GetFileExt((file.Name()))
		fmt.Printf("%s\t\t\t%s\t\t\t%d\t\t\t%s\n", file.Name(), finfo.GetFileType(ext), file.Size(), file.ModTime())
	}
}


func (finfo *FileInfo) GetFileExt(path string) string {
	return filepath.Ext(path)
}

func (finfo *FileInfo) GetFileType(ext string) string {
	imageExts := []string{".png", ".gif", ".jpg", ".jiff"}
	documentExts := []string{".txt", ".pdf", ".docx", ".ppt", ".pptx"}
	audioExts := []string{".mp3", ".aac", ".wav", ".flac"}
	videoExts := []string{".mp4", ".mkv", ".avi", ".flv"}
	fileType := ""

	if collections.ContainsString(imageExts, ext) {
		fileType = "Image"
	} else if collections.ContainsString(documentExts, ext) {
		fileType = "Document"
	} else if collections.ContainsString(audioExts, ext) {
		fileType = "Audio"
	} else if collections.ContainsString(documentExts, ext) {
		fileType = "Document"
	} else if collections.ContainsString(videoExts, ext) {
		fileType = "Video"
	} else {
		fileType = "Other"
	}
	return fileType
}


func (finfo *FileInfo) IsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	if fileInfo.IsDir() {
		return true
	} else {
		return false
	}
}


func (finfo *FileInfo) FileExists(folder string, fileName string) bool {
	filePath := folder + "/" + fileName
	if _, err := os.Stat(filePath); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}


func printChar(ch string, chSize int) {
	for i := 0; i < chSize; i++ {
		fmt.Print(ch)
	}
	fmt.Println()
}