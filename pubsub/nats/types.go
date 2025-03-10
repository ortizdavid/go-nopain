package pubsub

import "time"

type NatsConfig struct {
	URL           string
	User          string
	Password      string
	Timeout       time.Duration
	ReconnectWait time.Duration
	MaxReconnects int
}
