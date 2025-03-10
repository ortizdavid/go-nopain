package main

import (
	"log"
	"github.com/ortizdavid/go-nopain/_examples/pubsub/helpers"
	pubsub "github.com/ortizdavid/go-nopain/pubsub/rabbitmq"
)

func processMessageFromExchange() {

	consumer, _ := pubsub.NewRabbitMQConsumerDefault()

	err := pubsub.ProcessMessageFromExchange(consumer, pubsub.DefaultExchange("golang_exchange"), "golang_key", helpers.AddMessageToSlice)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	processMessageFromExchange()
}
