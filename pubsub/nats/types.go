package pubsub

import "time"

// NatsConfig defines the configuration settings for connecting to a NATS server
type NatsConfig struct {
	URL           string
	User          string
	Password      string
	Timeout       time.Duration
	ReconnectWait time.Duration
	MaxReconnects int
}
