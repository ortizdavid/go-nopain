package pubsub

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Consumer represents the  Consumer configuration.
type Consumer struct {
	Server Server
}

// NewConsumer creates a new Consumer instance with custom server configuration.
func NewConsumer(host string, port int, user string, password string) *Consumer {
	return &Consumer{
		Server: Server{
			Host:     host,
			Port:     port,
			User:     user,
			Password: password,
		},
	}
}

// NewConsumerDefault creates a new Consumer instance with default server configuration.
func NewConsumerDefault() *Consumer {
	return &Consumer{
		Server: Server{
			Host:     "localhost",
			Port:     5672,
			User:     "guest",
			Password: "guest",
		},
	}
}

// ConsumeFromQueue consumes messages from the specified queue.
// It establishes a connection to the  server, opens a channel, declares the queue,
// consumes messages from it, and logs the received messages continuously.
func (consumer *Consumer) ConsumeFromQueue(queue Queue) error {
	conn, err := amqp.Dial(serverURI(consumer.Server))
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
	// consume the messages
	msgs, err := consumeMessages(ch, q)
	if err != nil {
		return err
	}
	// Process received messages
	forever := make(chan bool)
	logMessages(err, msgs)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}

// ConsumeFromExchange consumes messages from the specified exchange with the given routing key.
// It establishes a connection to the  server, opens a channel, declares the exchange,
// creates a new queue, binds the queue to the exchange with the routing key,
// consumes messages from the queue, and logs the received messages continuously.
func (consumer *Consumer) ConsumeFromExchange(exchange Exchange, routingKey string) error {
	conn, err := amqp.Dial(serverURI(consumer.Server))
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
	// Create a new queue
	q, err := declareQueueDefault(ch)
	if err != nil {
		return err
	}
	// Bind the queue to the exchange with the routing key
	err = bindQueue(ch, q.Name, exchange, routingKey)
	if err != nil {
		return err
	}
	// Consume messages from the queue
	msgs, err := consumeMessages(ch, q)
	if err != nil {
		return err
	}
	// Process received messages
	forever := make(chan struct{})
	logMessages(err, msgs)
	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

	return nil
}

// ProcessMessageFromQueue consumes messages from the specified queue.
// It establishes a connection to the  server, opens a channel, declares the queue,
// and consumes messages from it. Each received message is processed using the provided function.
func ProcessMessageFromQueue[T any](consumer *Consumer, queue Queue, fn func(T) error) error {
	conn, err := amqp.Dial(serverURI(consumer.Server))
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
	// Consume messages from the queue
	msgs, err := consumeMessages(ch, q)
	if err != nil {
		return err
	}
	// Process received messages
	forever := make(chan struct{})
	processMessages(msgs, fn)
	<-forever

	return nil
}

// ProcessMessageFromExchange consumes messages from the specified exchange with the given routing key.
// It establishes a connection to the  server, opens a channel, declares the exchange,
// creates a new queue, binds the queue to the exchange with the routing key,
// and consumes messages from the queue. Each received message is processed using the provided function.
func ProcessMessageFromExchange[T any](consumer *Consumer, exchange Exchange, routingKey string, fn func(T) error) error {
	conn, err := amqp.Dial(serverURI(consumer.Server))
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
	// Create a new queue
	q, err := declareQueueDefault(ch)
	if err != nil {
		return err
	}
	// Bind the queue to the exchange with the routing key
	err = bindQueue(ch, q.Name, exchange, routingKey)
	if err != nil {
		return err
	}
	// Consume messages from the queue
	msgs, err := consumeMessages(ch, q)
	if err != nil {
		return err
	}
	// Process received messages
	forever := make(chan struct{})
	processMessages(msgs, fn)
	<-forever

	return nil
}
