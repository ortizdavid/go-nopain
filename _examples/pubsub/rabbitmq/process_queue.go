package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ortizdavid/go-nopain/filemanager"
	"github.com/ortizdavid/go-nopain/pubsub/rabbitmq"
)

type golangMessage struct {
	Text    string `json:"text"`
	Number  int    `json:"number"`
	Boolean bool   `json:"boolean"`
}

func addMessageToFile(msg golangMessage) error {
	var filemanager filemanager.FileManager
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	newContent := fmt.Sprintf("[%s] %s\n", currentTime, msg.Text)
	filemanager.WriteFile(".", "messages.txt", newContent)
	return nil
}

func processMessageFromQueue() {

	consumer, _ := pubsub.NewRabbitMQConsumerDefault()

	err := pubsub.ProcessMessageFromQueue(consumer, pubsub.DefaultQueue("golang_queue"), addMessageToFile)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	processMessageFromQueue()
}
