package pubsub

// ExchangeRMQ represents the configuration of a RabbitMQ exchange.
type ExchangeRMQ struct {
	Name string
	ExType ExchangeType
	Durable bool
	AutoDelete bool
	Internal bool
	NoWait bool
	Arguments map[string]interface{}
}

// ServerRMQ contains settings for connecting to the RabbitMQ server.
type ServerRMQ struct {
	Host string
	Port int
	User string
	Password string
}

// QueueRMQ contains settings for configuring RabbitMQ queues.
type QueueRMQ struct {
	Name string
	Durable bool
	Exclusive bool
	AutoDelete bool
	NoWait bool
	Arguments map[string]interface{}
}

// BindRMQ contains settings for configuring RabbitMQ bindings.
type BindRMQ struct {
	Queue      string
	Exchange   string
	RoutingKey string
	NoWait     bool
	Arguments  map[string]interface{}
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

