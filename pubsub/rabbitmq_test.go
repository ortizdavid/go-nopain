package pubsub

import (
	"testing"
)


type golangMessage struct {
	Text string `json:"text"`
    Number int `json:"number"`
    Boolean bool `json:"boolean"`
}


func TestPublisToQueueWithConfigs(t *testing.T) {
	message := golangMessage{
		Text:    "Message with configs",
		Number:  99,
		Boolean: false,
	}
	serverConf := ServerConfig{
		Host:     "127.0.0.1",
		Port:     5672,
		User:     "guest",
		Password: "guest",
	}

	rmq := NewRabbitMQClient(serverConf)
	rmq.QueueConfig = QueueConfig{
		Name:       "golang_queue",
		Durable:    false,
		Exclusive:  false,
		AutoDelete: false,
		NoWait:     false,
		Arguments:  nil,
	}

	if err := rmq.PublishToQueue(message); err != nil {
		t.Error(err)
	}
}


func TestPublishToQueueWithDefault(t *testing.T) {
	rmq := NewRabbitMQClientDefault("golang_queue")
	message := golangMessage{
		Text:    "Message with Default",
		Number:  1097,
		Boolean: true,
	}
	
	if err := rmq.PublishToQueue(message); err != nil {
		t.Error(err)
	}
}


func TestPublisToExchange(t *testing.T) {
	message := golangMessage{
		Text:    "Message to exchange",
		Number:  99,
		Boolean: false,
	}
	serverConf := ServerConfig{
		Host:     "127.0.0.1",
		Port:     5672,
		User:     "guest",
		Password: "guest",
	}

	rmq := NewRabbitMQClient(serverConf)
	rmq.ExchangeConfig = ExchangeConfig{
		Name:       "golang_exchange",
		ExType:     ExchangeFanout,
		Durable:    false,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Arguments:  map[string]interface{}{},
	}

	if err := rmq.PublishToExchange("", message); err != nil {
		t.Error(err)
	}
}