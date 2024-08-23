package main

import (
	"fmt"
	"log"
	"github.com/ortizdavid/go-nopain/pubsub"
)

type golangMessage struct {
	Text string `json:"text"`
    Number int `json:"number"`
    Boolean bool `json:"boolean"`
}

var msgSlice []golangMessage

func printMessage(msg golangMessage) error {
	fmt.Println(msg)
	return nil
}

func addMessageToSlice(msg golangMessage) error {
	msgSlice = append(msgSlice, msg)
	return nil
}

func processMessageFromQueue() {

	producer := pubsub.NewRabbitMQConsumerDefault()

	queue := pubsub.QueueRMQ{
		Name:       "golang_queue",
		Durable:    false,
		AutoDelete: false,
		Exclusive:   false,
		NoWait:     false,
		Arguments:  nil,
	}
	err := pubsub.ProcessMessageFromQueue(producer, queue, addMessageToSlice)
	if err != nil {
		log.Println(err)
	}

}

func main() {
	processMessageFromQueue()
	fmt.Println(len(msgSlice))
}

