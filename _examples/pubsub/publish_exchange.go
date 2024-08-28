package main

import (
	"log"

	"github.com/ortizdavid/go-nopain/pubsub"
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

	producer := pubsub.NewProducerDefault()

	exchange := pubsub.Exchange{
		Name:       "golang_exchange",
		ExType:     pubsub.ExchangeFanout,
		Durable:    false,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Arguments:  nil,
	}

	err := producer.PublishToExchange(exchange, "golang_key", message)
	if err != nil {
		log.Println(err)
	}
}

func main() {

	publishToExchange()
}
