package pubsub

import (
	"context"
	"fmt"
	"time"

	"github.com/ortizdavid/go-nopain/serialization"
	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQProducer represents the RabbitMQ producer configuration.
type RabbitMQProducer struct {
	ServerRMQ ServerRMQ
}

// NewRabbitMQProducer creates a new RabbitMQProducer instance with custom server configuration.
func NewRabbitMQProducer(host string, port int, user string, password string) RabbitMQProducer {
	return RabbitMQProducer{
		ServerRMQ: ServerRMQ{
			Host:     host,
			Port:     port,
			User:     user,
			Password: password,
		},
	}
}

// NewRabbitMQProducerDefault creates a new RabbitMQProducer instance with default server configuration.
func NewRabbitMQProducerDefault() RabbitMQProducer {
	return RabbitMQProducer{
		ServerRMQ: ServerRMQ{
			Host:     "localhost",
			Port:     5672,
			User:     "guest",
			Password: "guest",
		},
	}
}

// PublishToQueue publishes a message to a RabbitMQ queue.
func (rmq RabbitMQProducer) PublishToQueue(queue QueueRMQ, objMessage interface{}) error {
	// Establish connection to RabbitMQ
	conn, err := amqp.Dial(serverURI(rmq.ServerRMQ))
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %w", err)
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

// PublishToExchange publishes a message to a RabbitMQ exchange.
func (rmq RabbitMQProducer) PublishToExchange(exchange ExchangeRMQ, routingKey string, objMessage interface{}) error {
	// Establish connection to RabbitMQ
	conn, err := amqp.Dial(serverURI(rmq.ServerRMQ))
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %w", err)
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
