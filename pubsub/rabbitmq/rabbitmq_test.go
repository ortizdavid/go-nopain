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

	if err := rmq.PublishToQueue(DefaultQueue("golang_queue"), message); err != nil {
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

	if err := rmq.PublishToExchange(DefaultExchange("golang_exchange"), "golang_key", message); err != nil {
		t.Error(err)
	}
}
