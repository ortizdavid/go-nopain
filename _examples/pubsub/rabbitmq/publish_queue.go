package main

import (
	"log"

	"github.com/ortizdavid/go-nopain/pubsub/rabbitmq"
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

	producer, _ := pubsub.NewRabbitMQProducerDefault()

	err := producer.PublishToQueue(pubsub.DefaultQueue("golang_queue"), message)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	publishToQueue()
}
