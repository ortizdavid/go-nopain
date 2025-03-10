package main

import (
	"log"

	"github.com/ortizdavid/go-nopain/_examples/pubsub/helpers"
	pubsub "github.com/ortizdavid/go-nopain/pubsub/rabbitmq"
)

func processMessageFromQueue() {

	consumer, _ := pubsub.NewRabbitMQConsumerDefault()

	err := pubsub.ProcessMessageFromQueue(consumer, pubsub.DefaultQueue("golang_queue"), helpers.SaveMessageToFile)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	processMessageFromQueue()
}
