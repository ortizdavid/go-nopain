package main

import (
	"log"

	"github.com/ortizdavid/go-nopain/pubsub"
)



type golangMessage struct {
	Text string `json:"text"`
    Number int `json:"number"`
    Boolean bool `json:"boolean"`
}

func main() {

	//publishToQueue()
	publishToExchange()
}


func publishToQueue()  {
	message := golangMessage{
		Text:    "Nova sms",
		Number:  12,
		Boolean: false,
	}

	queue := pubsub.QueueRMQ{
		Name:       "golang_queue",
		Durable:    false,
		Exclusive:  false,
		AutoDelete: false,
		NoWait:     false,
		Arguments:  nil,
	}

	rmq := pubsub.NewRabbitMQProducerDefault()

	err := rmq.PublishToQueue(queue, message)
	if err != nil {
		log.Println(err)
	}
}


func publishToExchange()  {
	message := golangMessage{
		Text:    "Mensagem para Exchange",
		Number:  82736,
		Boolean: false,
	}

	rmq2 := pubsub.NewRabbitMQProducerDefault()

	exchange := pubsub.ExchangeRMQ{
		Name:       "golang_exchange",
		ExType:     pubsub.ExchangeFanout,
		Durable:    false,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Arguments:  nil,
	}

	err := rmq2.PublishToExchange(exchange, "golang_key", message)
	if err != nil {
		log.Println(err)
	}
}


