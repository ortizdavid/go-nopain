package pubsub

import (
	"fmt"
	"log"

	"github.com/ortizdavid/go-nopain/serialization"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Declares a queue on the given channel
func declareQueue(ch *amqp.Channel, queue QueueRMQ) (amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		queue.Name,
		queue.Durable,
		queue.AutoDelete,
		queue.Exclusive,
		queue.NoWait,
		amqp.Table(queue.Arguments),
	)
	if err != nil {
		return amqp.Queue{}, fmt.Errorf("failed to declare a queue: %w", err)
	}
	return q, nil
}


// Declares a queue on the given channel
func declareQueueDefault(ch *amqp.Channel) (amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		"",       // Name
		false,    // Durable
		false,    // Delete when unused
		true,     // Exclusive
		false,    // No-wait
		nil,      // Arguments
	)
	if err != nil {
		return amqp.Queue{}, fmt.Errorf("failed to declare a queue: %w", err)
	}
	return q, nil
}

// Declare exchange on the given channel
func declareExchange(ch *amqp.Channel, exchange ExchangeRMQ) error {
	err := ch.ExchangeDeclare(
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
	return nil
}

// Bind the queue to the exchange with the routing key
func bindQueue(ch *amqp.Channel, queueName string, exchange ExchangeRMQ, routingKey string) error {
	err := ch.QueueBind(
		queueName,             // Queue name
		routingKey,         // Routing key
		exchange.Name,      // Exchange name
		false,              // No-wait
		nil,                // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to bind the queue to the exchange: %w", err)
	}
	return nil
}

// Starts consuming messages from the specified queue
func consumeMessages(ch *amqp.Channel, q amqp.Queue) (<-chan amqp.Delivery, error) {
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
		return nil, fmt.Errorf("failed to register a consumer: %w", err)
	}
	return msgs, nil
}

// Consumes received messages
func logMessages(err error, msgs <-chan amqp.Delivery) {
	go func() {
		for msg := range msgs {
			if err != nil {
				fmt.Println(err)
				continue
			}
			log.Printf("Received a message: %s", msg.Body)
		}
	}()
}

// Process received message
func processMessages[T any](msgs <-chan amqp.Delivery, fn func(T) error) {
	go func() {
	  for d := range msgs {
		var message T
		err := serialization.UnserializeJson(d.Body, &message)
		if err != nil {
		  log.Printf("failed to unserialize message body: %s", err)
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
}

// serverURI returns the AMQP connection string.
func serverURI(server ServerRMQ) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/",
		server.User,
		server.Password,
		server.Host,
		server.Port)
}

