package main

import (
	"log"

	"github.com/ortizdavid/go-nopain/_examples/pubsub/helpers"
	pubsub "github.com/ortizdavid/go-nopain/pubsub/nats"
)

func main() {

	publisher, err := pubsub.NewNatsPublisherDefault()
	if err != nil {
		log.Println(err)
	}
	defer publisher.Close()

	message := helpers.GolangMessage{
		Text:    "Ortiz David",
		Number:  34,
		Boolean: true,
	}

	err = publisher.Publish("golang_test", message)
	if err != nil {
		log.Println(err)
	}
}
