package main

import (
	"log"
	"github.com/ortizdavid/go-nopain/_examples/pubsub/helpers"
	pubsub "github.com/ortizdavid/go-nopain/pubsub/rabbitmq"
)

func publishToExchange() {
	message := helpers.GolangMessage{
		Text:    "Message for para Exchange",
		Number:  82736,
		Boolean: false,
	}

	producer, _ := pubsub.NewRabbitMQProducerDefault()

	err := producer.PublishToExchange(pubsub.DefaultExchange("golang_exchange"), "golang_key", message)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	publishToExchange()
}
