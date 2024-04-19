package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQClient represents the RabbitMQ client configuration.
type RabbitMQClient struct {
	ServerConfig ServerConfig
	QueueConfig	QueueConfig
	ExchangeConfig ExchangeConfig
}

// ServerConfig contains settings for connecting to the RabbitMQ server.
type ServerConfig struct {
	Host string
	Port int
	User string
	Password string
}

// QueueConfig contains settings for configuring RabbitMQ queues.
type QueueConfig struct {
	Name string
	Durable bool
	Exclusive bool
	AutoDelete bool
	NoWait bool
	Arguments map[string]interface{}
}

// ExchangeConfig represents the configuration of a RabbitMQ exchange.
type ExchangeConfig struct {
	Name string
	ExType ExchangeType
	Durable bool
	AutoDelete bool
	Internal bool
	NoWait bool
	Arguments map[string]interface{}
}

// Type of exchanges
type ExchangeType string

// Constants defining various types of RabbitMQ exchanges.
const (
    ExchangeFanout ExchangeType = "fanout" 
    ExchangeDirect ExchangeType = "direct"
	ExchangeTopic ExchangeType = "topic"
	ExchangeHeaders ExchangeType = "headers"
)


func NewRabbitMQClient(serverConfig ServerConfig) RabbitMQClient {
	return RabbitMQClient{
		ServerConfig: serverConfig,
	}
}


func NewRabbitMQClientDefault(queueName string) RabbitMQClient {
	return RabbitMQClient{
		ServerConfig: ServerConfig{
			Host:     "localhost",
			Port:     5672,
			User:     "guest",
			Password: "guest",
		},
		QueueConfig:  QueueConfig{
			Name:       queueName,
			Durable:    false,
			Exclusive:  false,
			AutoDelete: false,
			NoWait:     false,
			Arguments:  nil,
		},
	}
}


// PublishToQueue publishes a message to a RabbitMQ queue.
func (rmq RabbitMQClient) PublishToQueue(objMessage interface{}) error {
    // Establish connection to RabbitMQ
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
        rmq.QueueConfig.Name,
        rmq.QueueConfig.Durable,
        rmq.QueueConfig.AutoDelete,
        rmq.QueueConfig.Exclusive,
        rmq.QueueConfig.NoWait,
        amqp.Table(rmq.QueueConfig.Arguments),
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
func (rmq RabbitMQClient) PublishToExchange(routingKey string, objMessage interface{}) error {
    // Establish connection to RabbitMQ
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
        rmq.ExchangeConfig.Name,                      // Exchange name
        string(rmq.ExchangeConfig.ExType), // Exchange type
        rmq.ExchangeConfig.Durable,        // Durable
        rmq.ExchangeConfig.AutoDelete,     // Auto-delete
        rmq.ExchangeConfig.Internal,       // Internal
        rmq.ExchangeConfig.NoWait,         // No-wait
        amqp.Table(rmq.ExchangeConfig.Arguments), // Arguments
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

    // Determine routing key (set to empty string for fanout exchanges)
    var key string
    if rmq.ExchangeConfig.ExType == ExchangeFanout {
        key = ""
    } else {
        key = routingKey
    }
  
    // Publish message to exchange
    err = ch.PublishWithContext(ctx, 
        rmq.ExchangeConfig.Name, // Exchange name
        key,           // Routing key (empty for fanout exchange)
        false,        // Mandatory
        false,        // Immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body, // Serialized message body
        })
    if err != nil {
        return fmt.Errorf("failed to publish message to exchange: %w", err)
    }

    return nil
}


func (rmq RabbitMQClient) connectionString() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", 
		rmq.ServerConfig.User,
		rmq.ServerConfig.Password,
		rmq.ServerConfig.Host,
		rmq.ServerConfig.Port)
}