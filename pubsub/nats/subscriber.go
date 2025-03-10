package pubsub

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

// NatsSubscriber represents a NATS subscriber instance
type NatsSubscriber struct {
	conn *nats.Conn
}

// NewNatsSubscriber creates a new NATS subscriber with the given configuration
func NewNatsSubscriber(config NatsConfig) (*NatsSubscriber, error) {
	opts := []nats.Option{
		nats.Timeout(config.Timeout),
		nats.MaxReconnects(config.MaxReconnects),
		nats.MaxReconnects(config.MaxReconnects),
	}

	if config.User != "" && config.Password != "" {
		opts = append(opts, nats.UserInfo(config.User, config.Password))
	}

	nc, err := nats.Connect(config.URL, opts...)
	if err != nil {
		return nil, err
	}

	return &NatsSubscriber{conn: nc}, nil
}

// NewNatsSubscriberDefault creates a subscriber with default configuration
func NewNatsSubscriberDefault() (*NatsSubscriber, error) {
	config := NatsConfig{
		URL:           "nats://localhost:4222",
		Timeout:       5 * time.Second,
		ReconnectWait: 2 * time.Second,
		MaxReconnects: 5,
	}
	return NewNatsSubscriber(config)
}

// Subscribe listens to messages on a given subject
func (sub *NatsSubscriber) Subscribe(subject string) error {
	fmt.Printf("Awaiting messages on subject '%s ...", subject)
	_, err := sub.conn.Subscribe(subject, func(msg *nats.Msg) {
		fmt.Printf("[x] Received message on '%s': %s\n", subject, string(msg.Data))
	})
	return err
}

// SubscribeQueue listens to messages on a subject using a queue group
func (sub *NatsSubscriber) SubscribeQueue(subject string, queue string) error {
	fmt.Printf("Awaiting messages on subject '%s, queue '%s' ...", subject, queue)
	_, err := sub.conn.QueueSubscribe(subject, queue, func(msg *nats.Msg) {
		fmt.Printf("[x] Received message on subject '%s', queue '%s': %s\n", subject, queue, string(msg.Data))
	})
	return err
}

// ProcessMessage processes incoming messages and applies a handler function
func ProcessMessage[T any](subscriber *NatsSubscriber, subject string, fn func(T) error) error {
	_, err := subscriber.conn.Subscribe(subject, func(msg *nats.Msg) {
		var message T
		err := json.Unmarshal(msg.Data, &message)
		if err != nil {
			fmt.Printf("[!] Error decoding message on '%s': %v", subject, err)
			return
		}
		err = fn(message)
		if err != nil {
			fmt.Printf("[!] Error processing message on '%s': %v", subject, err)
		}
	})
	return err
}

// ProcessMessageQueue processes messages from a queue and applies a handler function
func ProcessMessageQueue[T any](subscriber *NatsSubscriber, subject string, queue string, fn func(T) error) error {
	_, err := subscriber.conn.QueueSubscribe(subject, queue, func(msg *nats.Msg) {
		var message T
		err := json.Unmarshal(msg.Data, &message)
		if err != nil {
			fmt.Printf("[!] Error decoding message on '%s': %v", subject, err)
			return
		}
		err = fn(message)
		if err != nil {
			fmt.Printf("[!] Error processing message on '%s': %v", subject, err)
		}
	})
	return err
}

// Close terminates the NATS subscriber connection
func (pub *NatsSubscriber) Close() {
	if pub.conn != nil {
		pub.conn.Close()
	}
}