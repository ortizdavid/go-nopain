package main

import (
	"log"

	"github.com/ortizdavid/go-nopain/pubsub/rabbitmq"
)

func main() {
	consumeFromExchange()
}

func consumeFromExchange() {

	rmq2, _ := pubsub.NewRabbitMQConsumerDefault()

	err := rmq2.ConsumeFromExchange(pubsub.DefaultExchange("user.exchange"), "user.updated")
	if err != nil {
		log.Println(err)
	}
}
