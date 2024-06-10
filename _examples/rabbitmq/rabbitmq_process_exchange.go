package main

import (
	"fmt"
	"log"

	"github.com/ortizdavid/go-nopain/pubsub"
)

type golangMessage struct {
	Text string `json:"text"`
    Number int `json:"number"`
    Boolean bool `json:"boolean"`
}


func main() {
	processMessageFromExchange()
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
	err := pubsub.ProcessMessageFromExchange(rmq, exchange, "golang_key", addMessageToSlice)
	if err != nil {
		log.Println(err)
	}
}

func printMessage(msg golangMessage) error {
	fmt.Println(msg)
	return nil
}

var slices []golangMessage

func addMessageToSlice(msg golangMessage) error {
	slices = append(slices, msg)
	return nil
}

