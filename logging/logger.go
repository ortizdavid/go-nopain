package logging

import (
	"os"
	"log/slog"
)


func NewLogger(folderName string, fileName string) *slog.Logger {
	file, err := os.OpenFile(folderName+"/"+fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        panic(err)
    }
    defer file.Close()
	logger := slog.New(slog.NewJSONHandler(file, nil))
	return logger
}


