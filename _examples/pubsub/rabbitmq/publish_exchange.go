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

func publishToExchange() {
	message := golangMessage{
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
