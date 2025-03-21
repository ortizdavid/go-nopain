package pubsub

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

// NatsPublisher represents a NATS publisher for sending messages
type NatsPublisher struct {
	conn *nats.Conn
}

// NewNatsPublisher creates a new instance of NatsPublisher with the specified configuration
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

// NewNatsPublisherDefault creates a NATS publisher with default configurations
func NewNatsPublisherDefault() (*NatsPublisher, error) {
	config := NatsConfig{
		URL: "nats://localhost:4222",
		Timeout: 5 * time.Second,
		ReconnectWait: 2 * time.Second,
		MaxReconnects: 5,
	}
	return NewNatsPublisher(config)
}

// Publish sends a message to a specified subject in NATS
func (pub *NatsPublisher) Publish(subject string, message interface{}) error {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)

	if err := json.NewEncoder(gz).Encode(message); err != nil {
		gz.Close()
		return fmt.Errorf("[!] error encoding message: %v", err)
	}
	gz.Close()

	err := pub.conn.Publish(subject, buf.Bytes())
	if err != nil {
		return fmt.Errorf("error publishing to '%s': %v", subject, err)
	}
	return nil
}

// Close terminates the connection with the NATS server
func (pub *NatsPublisher) Close() {
	if pub.conn != nil {
		pub.conn.Close()
	}
}