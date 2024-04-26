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
// It establishes a connection to the RabbitMQ server, opens a channel, declares the queue,
// consumes messages from it, and logs the received messages continuously.
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
// It establishes a connection to the RabbitMQ server, opens a channel, declares the exchange,
// creates a new queue, binds the queue to the exchange with the routing key,
// consumes messages from the queue, and logs the received messages continuously.
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


// ProcessMessageFromQueue consumes messages from the specified queue.
// It establishes a connection to the RabbitMQ server, opens a channel, declares the queue,
// and consumes messages from it. Each received message is processed using the provided function.
func ProcessMessageFromQueue[T any](rmq RabbitMQConsumer, queue QueueRMQ, fn func(T) error) error {
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
	forever := make(chan struct{})
	processMessages(msgs, fn)
	<-forever
  
	return nil
}


// ProcessMessageFromExchange consumes messages from the specified exchange with the given routing key.
// It establishes a connection to the RabbitMQ server, opens a channel, declares the exchange,
// creates a new queue, binds the queue to the exchange with the routing key,
// and consumes messages from the queue. Each received message is processed using the provided function.
func ProcessMessageFromExchange[T any](rmq RabbitMQConsumer, exchange ExchangeRMQ, routingKey string, fn func(T) error) error {
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


