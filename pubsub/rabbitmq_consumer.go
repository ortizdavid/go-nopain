package pubsub

import (
	"fmt"
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQConsumer represents the RabbitMQ Consumer configuration.
type RabbitMQConsumer struct {
	ServerRMQ ServerRMQ
}

// NewRabbitMQConsumer creates a new RabbitMQConsumer instance with custom server configuration.
func NewRabbitMQConsumer(host string, port int, user string, password string) RabbitMQConsumer {
	return RabbitMQConsumer{
		ServerRMQ: ServerRMQ{
			Host:     host,
			Port:     port,
			User:     user,
			Password: password,
		},
	}
}


// NewRabbitMQConsumerDefault creates a new RabbitMQConsumer instance with default server configuration.
func NewRabbitMQConsumerDefault() RabbitMQConsumer {
	return RabbitMQConsumer{
		ServerRMQ: ServerRMQ{
			Host:     "localhost",
			Port:     5672,
			User:     "guest",
			Password: "guest",
		},
	}
}


// ConsumeFromQueue consumes messages from the specified queue.
func (rmq RabbitMQConsumer) ConsumeFromQueue(queue QueueRMQ) error {
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
func (rmq RabbitMQConsumer) ConsumeFromExchange(exchange ExchangeRMQ, routingKey string) error {
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


// ProcessMessageFromQueue consumes and processes messages from the specified queue.
func (rmq RabbitMQConsumer) ProcessMessageFromQueue(queue QueueRMQ, fn func(message interface{}) error) error {
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

	// Consume messages from the queue
	msgs, err := consumeMessages(ch, q)
	if err != nil {
		return err
	}

	// Process received messages
	forever := make(chan bool)
	processMessages(msgs, fn)
	<-forever

	return nil
}

// ConsumeFromExchange consumes messages from the specified exchange with the given routing key.
func (rmq RabbitMQConsumer) ProcessMessageFromExchange(exchange ExchangeRMQ, routingKey string, fn func(message interface{}) error) error {
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
