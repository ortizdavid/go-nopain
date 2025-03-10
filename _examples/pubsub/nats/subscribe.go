package main

import (
	"log"

	pubsub "github.com/ortizdavid/go-nopain/pubsub/nats"
)

func main() {

	subscriber, err := pubsub.NewNatsSubscriberDefault()
	if err != nil {
		log.Println(err)
	}
	defer subscriber.Close()

	err = subscriber.Subscribe("golang_test")
	if err != nil {
		log.Println(err)
	}

	select {}
}