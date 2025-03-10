package main

import (
	"log"
	"github.com/ortizdavid/go-nopain/pubsub/rabbitmq"
)

func main() {
	rmq2, _ := pubsub.NewRabbitMQConsumerDefault()

	err := rmq2.ConsumeFromQueue(pubsub.DefaultQueue("golang_queue"))
	if err != nil {
		log.Println(err)
	}
}
