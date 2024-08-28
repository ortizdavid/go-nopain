package pubsub

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

// Server contains settings for connecting to the  server.
type Server struct {
	Host     string
	Port     int
	User     string
	Password string
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
