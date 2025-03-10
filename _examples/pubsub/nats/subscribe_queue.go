package main

import (
	"log"

	pubsub "github.com/ortizdavid/go-nopain/pubsub/nats"
)

func main()  {
	subscriber, err := pubsub.NewNatsSubscriberDefault()
	if err != nil {
		log.Println(err)
	}
	defer subscriber.Close()

	err = subscriber.SubscribeQueue("golang_test", "test.queue")
	if err != nil {
		log.Println(err)
	}

	select {}

}