package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQClient struct {
	ServerConfig ServerConfig
	QueueConfig	QueueConfig
}

type ServerConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

type QueueConfig struct {
	Name       string
	Durable    bool
	Exclusive  bool
	AutoDelete bool
	NoWait     bool
	Arguments map[string]interface{}
}


func NewRabbitMQClient(serverConfig ServerConfig, queueConfig QueueConfig) *RabbitMQClient {
	return &RabbitMQClient{
		ServerConfig: serverConfig,
		QueueConfig:  queueConfig,
	}
}


func NewRabbitMQClientDefault(queueName string) *RabbitMQClient {
	return &RabbitMQClient{
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


func (rmq *RabbitMQClient) PublishMessage(objMessage interface{}) error {
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

    // Publish message
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

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


func (rmq *RabbitMQClient) connectionString() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", 
		rmq.ServerConfig.User,
		rmq.ServerConfig.Password,
		rmq.ServerConfig.Host,
		rmq.ServerConfig.Port)
}