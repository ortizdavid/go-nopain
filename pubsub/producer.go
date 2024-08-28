package pubsub

import (
	"context"
	"fmt"
	"time"

	"github.com/ortizdavid/go-nopain/serialization"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Producer represents the  producer configuration.
type Producer struct {
	config ServerConfig
}

// NewProducer creates a new Producer instance with custom server configuration.
func NewProducer(config ServerConfig) *Producer {
	// Basic config
	if config.Host == "" {
		panic("host cannot be empty")
	}
	if config.Port <= 0 || config.Port > 65535 {
		panic("invalid port number")
	}
	if config.User == "" {
		panic("user cannot be empty")
	}
	if config.Password == "" {
		panic("password cannot be empty")
	}
	return &Producer{
		config: config,
	}
}

// NewProducerDefault creates a new Producer instance with default server configuration.
func NewProducerDefault() *Producer {
	return &Producer{
		config: ServerConfig{
			Host:     "localhost",
			Port:     5672,
			User:     "guest",
			Password: "guest",
		},
	}
}

// PublishToQueue publishes a message to a  queue.
func (producer *Producer) PublishToQueue(queue Queue, objMessage interface{}) error {
	// Establish connection to
	conn, err := amqp.Dial(serverURI(producer.config))
	if err != nil {
		return fmt.Errorf("failed to connect to : %w", err)
	}
	defer conn.Close()
	// Open channel
	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %w", err)
	}
	defer ch.Close()
	// Declare queue
	q, err := declareQueue(ch, queue)
	if err != nil {
		return err
	}
	// Marshal message to JSON
	body, err := serialization.SerializeJson(objMessage)
	if err != nil {
		return fmt.Errorf("failed to serialize message to JSON: %w", err)
	}
	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Publish message
	err = ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}

// PublishToExchange publishes a message to a  exchange.
func (producer *Producer) PublishToExchange(exchange Exchange, routingKey string, objMessage interface{}) error {
	// Establish connection to
	conn, err := amqp.Dial(serverURI(producer.config))
	if err != nil {
		return fmt.Errorf("failed to connect to : %w", err)
	}
	defer conn.Close()
	// Open channel
	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %w", err)
	}
	defer ch.Close()
	// Declare exchange
	err = declareExchange(ch, exchange)
	if err != nil {
		return err
	}
	// Serialize message to JSON
	body, err := serialization.SerializeJson(objMessage)
	if err != nil {
		return fmt.Errorf("failed to serialize message to JSON: %w", err)
	}
	// Create Context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Determine routing key
	var key string
	if exchange.ExType == ExchangeFanout {
		key = ""
	} else {
		key = routingKey
	}
	// Publish message to exchange
	err = ch.PublishWithContext(ctx,
		exchange.Name,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return fmt.Errorf("failed to publish message to exchange: %w", err)
	}

	return nil
}
