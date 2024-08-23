package main


import (
	"fmt"
	"log"

	"github.com/ortizdavid/go-nopain/pubsub"
)


func main() {
	consumeFromExchange()
}

func consumeFromExchange() {

	rmq2 := pubsub.NewRabbitMQConsumerDefault()

	exchange := pubsub.ExchangeRMQ{
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
