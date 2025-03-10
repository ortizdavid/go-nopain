package helpers

import (
	"fmt"
	"time"

	"github.com/ortizdavid/go-nopain/filemanager"
)

type GolangMessage struct {
	Text    string `json:"text"`
	Number  int    `json:"number"`
	Boolean bool   `json:"boolean"`
}

var slices []GolangMessage

func PrintMessage(msg GolangMessage) error {
	fmt.Println(msg)
	return nil
}

func AddMessageToSlice(msg GolangMessage) error {
	slices = append(slices, msg)
	return nil
}

func SaveMessageToFile(msg GolangMessage) error {
	var filemanager filemanager.FileManager
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	newContent := fmt.Sprintf("[%s] %s\n", currentTime, msg.Text)
	filemanager.WriteFile(".", "../helpers/messages.txt", newContent)
	return nil
}
