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

	queue := pubsub.Queue{
		Name:       "golang_queue",
		Durable:    false,
		Exclusive:  false,
		AutoDelete: false,
		NoWait:     false,
		Arguments:  nil,
	}

	err := rmq2.ConsumeFromQueue(queue)
	if err != nil {
		log.Println(err)
	}

}
