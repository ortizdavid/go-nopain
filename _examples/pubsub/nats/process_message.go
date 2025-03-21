package main

import (
	"fmt"
	"log"

	"github.com/ortizdavid/go-nopain/_examples/pubsub/helpers"
	pubsub "github.com/ortizdavid/go-nopain/pubsub/nats"
	"github.com/ortizdavid/go-nopain/reflection"
)

func main() {
	subscriber, err := pubsub.NewNatsSubscriberDefault()
	if err != nil {
		log.Println(err)
	}
	defer subscriber.Close()

	err = pubsub.ProcessMessage(subscriber, "golang_test", helpers.SaveMessageToFile)
	if err != nil {
		log.Println(err)
	}
	select {}
	
}
