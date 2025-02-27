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

	err := rmq2.ConsumeFromQueue(pubsub.DefaultQueue("dotnet_queue"))
	if err != nil {
		log.Println(err)
	}

}
