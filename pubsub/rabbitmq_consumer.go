package pubsub

import (
	"fmt"
	"log"

	"github.com/ortizdavid/go-nopain/serialization"
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
func (rmq RabbitMQConsumer) ConsumeFromExchange(exchange ExchangeRMQ, routingKey string) error {
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


// ProcessMessageFromQueue consumes and processes messages from the specified queue.
func (rmq RabbitMQConsumer) ProcessMessageFromQueue(queue QueueRMQ, fn func(message interface{}) error) error {
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
			var message interface{}
			err := serialization.UnserializeJson(msg.Body, &message)
			if err != nil {
				log.Printf("failed to serialize message body: %s", err)
				continue
			}
			err = fn(message)
			if err != nil {
				log.Printf("error processing message: %s", err)
				continue
			}
			log.Printf("Processed message: %+v", message)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}

// ConsumeFromExchange consumes messages from the specified exchange with the given routing key.
func (rmq RabbitMQConsumer) ProcessMessageFromExchange(exchange ExchangeRMQ, routingKey string, fn func(message interface{}) error) error {
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
			var message interface{}
			err := serialization.UnserializeJson(d.Body, &message)
			if err != nil {
				log.Printf("failed to userialize message body: %s", err)
				continue
			}
			err = fn(message)
			if err != nil {
				log.Printf("error processing message: %s", err)
				continue
			}
			log.Printf("Processed message: %+v", message)
		}
	}()

	<-forever

	return nil
}


// serverURI returns the AMQP connection string.
func (rmq RabbitMQConsumer) serverURI() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/",
		rmq.ServerRMQ.User,
		rmq.ServerRMQ.Password,
		rmq.ServerRMQ.Host,
		rmq.ServerRMQ.Port)
}
