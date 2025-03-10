package pubsub

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

type NatsPublisher struct {
	conn *nats.Conn
}

func NewNatsPublisher(config NatsConfig) (*NatsPublisher, error) {
	opts := []nats.Option{
		nats.Timeout(config.Timeout),
		nats.ReconnectWait(config.ReconnectWait),
		nats.MaxReconnects(config.MaxReconnects),
	}

	if config.User != "" && config.Password != "" {
		opts = append(opts, nats.UserInfo(config.User, config.Password))
	}

	nc, err := nats.Connect(config.URL, opts...)
	if err != nil {
		return nil, err
	}

	return &NatsPublisher{conn: nc}, nil
}

func NewNatsPublisherDefault() (*NatsPublisher, error) {
	config := NatsConfig{
		URL: "nats://localhost:4222",
		Timeout: 5 * time.Second,
		ReconnectWait: 2 * time.Second,
		MaxReconnects: 5,
	}
	return NewNatsPublisher(config)
}

func (pub *NatsPublisher) Publish(subject string, message interface{}) error {
	data, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("[!] error encoding message: %v", err)
	}
	err = pub.conn.Publish(subject, data)
	if err != nil {
		return fmt.Errorf("error publishing to '%s': %v", subject, err)
	}
	return nil
}


func (pub *NatsPublisher) Close() {
	if pub.conn != nil {
		pub.conn.Close()
	}
}