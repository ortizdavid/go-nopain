package main

import (
	"fmt"
	"log"

	"github.com/ortizdavid/go-nopain/pubsub"
)


func main()  {
	//consumeFromExchange()
	processMessageFromExchange()
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


func processMessageFromExchange() {

	rmq := pubsub.NewRabbitMQConsumerDefault()

	exchange := pubsub.ExchangeRMQ{
		Name:       "golang_exchange",
		ExType:     pubsub.ExchangeFanout,
		Durable:    false,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Arguments:  nil,
	}

	err := pubsub.ProcessMessageFromExchange(rmq, exchange, "golang_key", printMessage)
	if err != nil {
		log.Println(err)
	}

}


func printMessage(msg golangMessage) error {
	fmt.Println(msg)
	return nil
}


type golangMessage struct {
	Text string `json:"text"`
    Number int `json:"number"`
    Boolean bool `json:"boolean"`
}