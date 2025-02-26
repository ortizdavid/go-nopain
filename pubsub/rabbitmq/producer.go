package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQProducer represents the  producer configuration.
type RabbitMQProducer struct {
	config     RabbitMQConfig
	connection *amqp.Connection
	channel    *amqp.Channel
}

// NewRabbitMQProducer creates a new RabbitMQProducer instance with custom server configuration.
func NewRabbitMQProducer(config RabbitMQConfig) (*RabbitMQProducer, error) {
	err := validateConfig(config)
	if err != nil {
		return nil, err
	}
	// Connection
	conn, err := amqp.Dial(connectionString(config))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to : %w", err)
	}
	// Open channel
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}
	return &RabbitMQProducer{
		config:     config,
		connection: conn,
		channel:    ch,
	}, nil
}

// NewRabbitMQProducerDefault creates a new RabbitMQProducer instance with default server configuration.
func NewRabbitMQProducerDefault() (*RabbitMQProducer, error) {
	return NewRabbitMQProducer(DefaultRabbitMQConfig)
}

// Close connection and Channel
func (producer *RabbitMQProducer) Close() {
	if producer.connection != nil {
		producer.connection.Close()
	}
	if producer.channel != nil {
		producer.channel.Close()
	}
}

// PublishToQueue publishes a message to a  queue.
func (producer *RabbitMQProducer) PublishToQueue(queue Queue, objMessage interface{}) error {
	ch := producer.channel
	// Declare queue
	q, err := declareQueue(ch, queue)
	if err != nil {
		return err
	}
	// Marshal message to JSON
	body, err := json.Marshal(objMessage)
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
func (producer *RabbitMQProducer) PublishToExchange(exchange Exchange, routingKey string, objMessage interface{}) error {
	ch := producer.channel
	// Declare exchange
	err := declareExchange(ch, exchange)
	if err != nil {
		return err
	}
	// Serialize message to JSON
	body, err := json.Marshal(objMessage)
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
