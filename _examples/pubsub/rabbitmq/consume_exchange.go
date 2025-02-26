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

	exchange := pubsub.Exchange{
		Name:       "golang_exchange",
		ExType:     pubsub.ExchangeFanout,
		Durable:    false,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Arguments:  nil,
	}

	err := rmq2.ConsumeFromExchange(exchange, "golang_key")
	if err != nil {
		log.Println(err)
	}

}
