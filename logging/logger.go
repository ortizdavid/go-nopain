package logging

import (
	"os"
	"log/slog"
)


// NewLogger creates a new logger with the specified folderName and fileName.
// It returns a pointer to the created logger.
func NewLogger(folderName string, fileName string) *slog.Logger {
	file, err := os.OpenFile(folderName+"/"+fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        panic(err)
    }
    defer file.Close()
	logger := slog.New(slog.NewJSONHandler(file, nil))
	return logger
}


