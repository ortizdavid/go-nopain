package pubsub

import (
	"testing"
)

type golangMessage struct {
	Text    string `json:"text"`
	Number  int    `json:"number"`
	Boolean bool   `json:"boolean"`
}

func TestPublishToQueueWithDefault(t *testing.T) {
	rmq, _ := NewRabbitMQProducerDefault()

	message := golangMessage{
		Text:    "Message with Default",
		Number:  1097,
		Boolean: true,
	}

	queue := Queue{
		Name:       "golang_queue",
		Durable:    false,
		Exclusive:  false,
		AutoDelete: false,
		NoWait:     false,
		Arguments:  nil,
	}

	if err := rmq.PublishToQueue(queue, message); err != nil {
		t.Error(err)
	}
}

func TestPublishToExchangeDefault(t *testing.T) {
	rmq, _ := NewRabbitMQProducerDefault()

	message := golangMessage{
		Text:    "Message to exchange",
		Number:  99,
		Boolean: false,
	}

	exchange := Exchange{
		Name:       "golang_exchange",
		ExType:     ExchangeFanout,
		Durable:    false,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Arguments:  map[string]interface{}{},
	}

	if err := rmq.PublishToExchange(exchange, "golang_key", message); err != nil {
		t.Error(err)
	}
}
