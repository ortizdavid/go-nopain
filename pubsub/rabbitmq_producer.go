package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
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
	conn, err := amqp.Dial(rmq.serverURI())
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
	q, err := ch.QueueDeclare(
		queue.Name,
		queue.Durable,
		queue.AutoDelete,
		queue.Exclusive,
		queue.NoWait,
		amqp.Table(queue.Arguments),
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %w", err)
	}

	// Marshal message to JSON
	body, err := json.Marshal(objMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal message to JSON: %w", err)
	}

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
	conn, err := amqp.Dial(rmq.serverURI())
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
	err = ch.ExchangeDeclare(
		exchange.Name,
		string(exchange.ExType),
		exchange.Durable,
		exchange.AutoDelete,
		exchange.Internal,
		exchange.NoWait,
		amqp.Table(exchange.Arguments),
	)
	if err != nil {
		return fmt.Errorf("failed to declare exchange: %w", err)
	}

	// Marshal message to JSON
	body, err := json.Marshal(objMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal message to JSON: %w", err)
	}

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

// serverURI returns the AMQP connection string.
func (rmq RabbitMQProducer) serverURI() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/",
		rmq.ServerRMQ.User,
		rmq.ServerRMQ.Password,
		rmq.ServerRMQ.Host,
		rmq.ServerRMQ.Port)
}
