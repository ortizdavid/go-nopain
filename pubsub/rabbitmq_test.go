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
	queueConf := QueueConfig{
		Name:       "golang_queue",
		Durable:    false,
		Exclusive:  false,
		AutoDelete: false,
		NoWait:     false,
		Arguments:  nil,
	}
	
	rmq := NewRabbitMQClient(serverConf, queueConf)

	if err := rmq.PublishMessage(message); err != nil {
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
	
	if err := rmq.PublishMessage(message); err != nil {
		t.Error(err)
	}
}