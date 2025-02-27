package pubsub

// RabbitMQConfig contains settings for connecting to the  server.
type RabbitMQConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

// Exchange represents the configuration of a  exchange.
type Exchange struct {
	Name       string
	ExType     ExchangeType
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Arguments  map[string]interface{}
}

// Queue contains settings for configuring  queues.
type Queue struct {
	Name       string
	Durable    bool
	Exclusive  bool
	AutoDelete bool
	NoWait     bool
	Arguments  map[string]interface{}
}

// Bind contains settings for configuring  bindings.
type Bind struct {
	Queue      string
	Exchange   string
	RoutingKey string
	NoWait     bool
	Arguments  map[string]interface{}
}

// Type of exchanges
type ExchangeType string

// Constants defining various types of  exchanges.
const (
	ExchangeFanout  ExchangeType = "fanout"
	ExchangeDirect  ExchangeType = "direct"
	ExchangeTopic   ExchangeType = "topic"
	ExchangeHeaders ExchangeType = "headers"
)

// Default config
var DefaultRabbitMQConfig = RabbitMQConfig{
	Host:     "localhost",
	Port:     5672,
	User:     "guest",
	Password: "guest",
}

// Default queue
var DefaultQueue = func(name string) Queue{
	return Queue{
		Name: name,
		Durable: false,
		Exclusive: false,
		AutoDelete: false,
		NoWait: false,
		Arguments: nil,
	}
} 

// Default exchange
var DefaultExchange = func(name string) Exchange{
	return Exchange{
		Name: name,
		ExType: ExchangeDirect,
		Durable: false,
		AutoDelete: false,
		Internal: false,
		NoWait: false,
		Arguments: nil,
	}
} 
