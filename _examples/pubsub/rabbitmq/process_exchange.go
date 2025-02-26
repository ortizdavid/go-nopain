package main

import (
	"fmt"
	"log"

	"github.com/ortizdavid/go-nopain/pubsub/rabbitmq"
)

type golangMessage struct {
	Text    string `json:"text"`
	Number  int    `json:"number"`
	Boolean bool   `json:"boolean"`
}

var slices []golangMessage

func printMessage(msg golangMessage) error {
	fmt.Println(msg)
	return nil
}

func addMessageToSlice(msg golangMessage) error {
	slices = append(slices, msg)
	return nil
}

func processMessageFromExchange() {

	consumer, _ := pubsub.NewRabbitMQConsumerDefault()
	exchange := pubsub.Exchange{
		Name:       "golang_exchange",
		ExType:     pubsub.ExchangeFanout,
		Durable:    false,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Arguments:  nil,
	}
	err := pubsub.ProcessMessageFromExchange(*consumer, exchange, "golang_key", addMessageToSlice)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	processMessageFromExchange()
}
