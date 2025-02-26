package pubsub

import (
	"fmt"
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQConsumer represents the  RabbitMQConsumer configuration.
type RabbitMQConsumer struct {
	config     RabbitMQConfig
	connection *amqp.Connection
	channel    *amqp.Channel
}

// NewRabbitMQConsumer creates a new RabbitMQConsumer instance with custom server configuration.
func NewRabbitMQConsumer(config RabbitMQConfig) (*RabbitMQConsumer, error) {
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
	return &RabbitMQConsumer{
		config:     config,
		connection: conn,
		channel:    ch,
	}, nil
}

// NewRabbitMQConsumerDefault creates a new RabbitMQConsumer instance with default server configuration.
func NewRabbitMQConsumerDefault() (*RabbitMQConsumer, error) {
	return NewRabbitMQConsumer(DefaultRabbitMQConfig)
}

// Close connection and Channel
func (consumer *RabbitMQConsumer) Close() {
	if consumer.connection != nil {
		consumer.connection.Close()
	}
	if consumer.channel != nil {
		consumer.channel.Close()
	}
}

// ConsumeFromQueue consumes messages from the specified queue.
// It consumes messages from it, and logs the received messages continuously.
func (consumer *RabbitMQConsumer) ConsumeFromQueue(queue Queue) error {
	ch := consumer.channel
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
// It consumes messages from the queue, and logs the received messages continuously.
func (consumer *RabbitMQConsumer) ConsumeFromExchange(exchange Exchange, routingKey string) error {
	ch := consumer.channel
	// Declare exchange
	err := declareExchange(ch, exchange)
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
// Each received message is processed using the provided function.
func ProcessMessageFromQueue[T any](consumer *RabbitMQConsumer, queue Queue, fn func(T) error) error {
	ch := consumer.channel
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
// It consumes messages from the queue. Each received message is processed using the provided function.
func ProcessMessageFromExchange[T any](consumer *RabbitMQConsumer, exchange Exchange, routingKey string, fn func(T) error) error {
	ch := consumer.channel
	// Declare exchange
	err := declareExchange(ch, exchange)
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
