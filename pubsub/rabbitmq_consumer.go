package pubsub

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQConsumer represents the RabbitMQ Consumer configuration.
type RabbitMQConsumer struct {
	ServerConfig ServerConfig
}

// NewRabbitMQConsumer creates a new RabbitMQConsumer instance with custom server configuration.
func NewRabbitMQConsumer(host string, port int, user string, password string) RabbitMQConsumer {
	return RabbitMQConsumer{
		ServerConfig: ServerConfig{
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
		ServerConfig: ServerConfig{
			Host:     "localhost",
			Port:     5672,
			User:     "guest",
			Password: "guest",
		},
	}
}

// ConsumeFromQueue consumes messages from the specified queue.
func (rmq RabbitMQConsumer) ConsumeFromQueue(queue QueueConfig) error {
	conn, err := amqp.Dial(rmq.connectionString())
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

	// Consume messages from the queue
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %w", err)
	}

	// Process received messages
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			if err != nil {
				fmt.Println(err)
				continue
			}
			log.Printf("Received a message: %s", msg.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}

// ConsumeFromExchange consumes messages from the specified exchange with the given routing key.
func (rmq RabbitMQConsumer) ConsumeFromExchange(exchange ExchangeConfig, routingKey string) error {
	conn, err := amqp.Dial(rmq.connectionString())
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
		exchange.Name,                      // Exchange name
		string(exchange.ExType),            // Exchange type
		exchange.Durable,                   // Durable
		exchange.AutoDelete,                // Auto-delete
		exchange.Internal,                  // Internal
		exchange.NoWait,                    // No-wait
		amqp.Table(exchange.Arguments),    // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare an exchange: %w", err)
	}

	// Create a new queue
	q, err := ch.QueueDeclare(
		"",       // Name
		false,    // Durable
		false,    // Delete when unused
		true,     // Exclusive
		false,    // No-wait
		nil,      // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %w", err)
	}

	// Bind the queue to the exchange with the routing key
	err = ch.QueueBind(
		q.Name,             // Queue name
		routingKey,         // Routing key
		exchange.Name,      // Exchange name
		false,              // No-wait
		nil,                // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to bind the queue to the exchange: %w", err)
	}

	// Consume messages from the queue
	msgs, err := ch.Consume(
		q.Name, // Queue name
		"",     // Consumer
		true,   // Auto-acknowledge
		false,  // Exclusive
		false,  // No-local
		false,  // No-wait
		nil,    // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %w", err)
	}

	// Process received messages
	forever := make(chan struct{})
	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

	return nil
}

// connectionString returns the AMQP connection string.
func (rmq RabbitMQConsumer) connectionString() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/",
		rmq.ServerConfig.User,
		rmq.ServerConfig.Password,
		rmq.ServerConfig.Host,
		rmq.ServerConfig.Port)
}
