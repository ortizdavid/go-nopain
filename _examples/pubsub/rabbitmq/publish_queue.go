package main

import (
	"log"

	"github.com/ortizdavid/go-nopain/_examples/pubsub/helpers"
	pubsub "github.com/ortizdavid/go-nopain/pubsub/rabbitmq"
)

func main() {
	message := helpers.GolangMessage{
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
