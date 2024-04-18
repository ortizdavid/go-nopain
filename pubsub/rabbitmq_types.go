package pubsub

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