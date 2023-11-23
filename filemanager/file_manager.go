package filemanager

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

type FileManager struct {}

func (fm *FileManager) CreateSingleFolder(folderName string) bool {
	err := os.Mkdir(folderName, 0777)
	if err == nil {
		log.Printf(FOLDER_CREATED, folderName)
		return true
	}
	log.Fatal(err.Error())
	return false
}


func (fm *FileManager) CreateManyFolders(folderName string) bool {
	err := os.MkdirAll(folderName, 0777)
	if err == nil {
		log.Printf(FOLDER_CREATED, folderName)
		return true
	} 
	log.Fatal(err.Error())
	return false
}


func (fm *FileManager) CreateSingleFile(dirName string, fileName string) bool {
	file, err := os.Create(dirName +"/"+ fileName)
	if err == nil {
		log.Printf(FILE_CREATED, file.Name())
		return true
	}
	log.Fatal(err.Error())
	return false
}


func (fm *FileManager) CreateManyFiles(dirName string, files ...string) bool {
	for _, file := range files {
		created := fm.CreateSingleFile(dirName, file)
		if created == false {
			return false
		}
	}
	return true
}


func (fm *FileManager) MoveFile(fileName string, origin string, destination string) bool {
	err := os.Rename(origin +"/"+ fileName, destination +"/"+ fileName)
	if err == nil {
		return true
	}
	log.Fatal(err.Error())
	return true
}


func (fm *FileManager) WriteFile(folderName string, fileName string, content string) bool {
	file, err := os.OpenFile(folderName +"/"+ fileName, os.O_APPEND, 0666)
	if err == nil {
		file.WriteString(content)
		file.Close()
		return true
	}
	log.Fatal(err.Error())
	return false
}


func (fm *FileManager) RemoveFile(folderName string, fileName string) bool {
	err := os.Remove(folderName+"/"+fileName)
	if err == nil {
		log.Printf(FILE_REMOVED, folderName+"/"+fileName)
		return true
	}
	log.Fatal(err.Error())
	return false
}


func (fm *FileManager) RemoveFolder(folderName string) bool {
	err := os.RemoveAll(folderName)
	if err == nil {
		log.Printf(FOLDER_REMOVED, folderName)
		return true
	}
	log.Fatal(err.Error())
	return false
}


func (fm *FileManager) CopyFolder(src string, dest string) error {
	err := os.MkdirAll(dest, os.ModePerm)
	if err != nil {
		return err
	}
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		if entry.IsDir() {
			if err := fm.CopyFolder(srcPath, destPath); err != nil {
				return err
			}
		} else {
			if err := fm.CopyFile(srcPath, destPath); err != nil {
				return err
			}
		}
	}
	return nil
}


func (fm *FileManager) CopyFile(src string, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}
	return nil
}

