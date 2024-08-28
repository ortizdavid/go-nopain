package main

import (
	"log"

	"github.com/ortizdavid/go-nopain/pubsub"
)

type golangMessage struct {
	Text    string `json:"text"`
	Number  int    `json:"number"`
	Boolean bool   `json:"boolean"`
}

func publishToQueue() {

	message := golangMessage{
		Text:    "New message for queue",
		Number:  12,
		Boolean: false,
	}

	queue := pubsub.Queue{
		Name:       "golang_queue",
		Durable:    false,
		Exclusive:  false,
		AutoDelete: false,
		NoWait:     false,
		Arguments:  nil,
	}

	producer := pubsub.NewProducerDefault()

	err := producer.PublishToQueue(queue, message)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	publishToQueue()
}
