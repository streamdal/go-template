package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

const (
	EnvConfigPrefix = "GO_TEMPLATE"
)

type Config struct {
	ListenAddress   string `envconfig:"LISTEN_ADDRESS" default:":8282"`
	HealthFreqSec   int    `envconfig:"HEALTH_FREQ_SEC" default:"60"`
	EnvName         string `envconfig:"ENV_NAME" default:"local"`
	ServiceName     string `envconfig:"SERVICE_NAME" default:"go-template"`
	NewRelicLicense string `envconfig:"NEW_RELIC_LICENSE"`

	// Queue for _internal_ events
	EventBusRabbitURL          string `envconfig:"EVENT_BUS_RABBIT_URL" default:"amqp://localhost"`
	EventBusRabbitExchangeName string `envconfig:"EVENT_BUS_RABBIT_EXCHANGE_NAME" default:"events"`
	EventBusRabbitRoutingKey   string `envconfig:"EVENT_BUS_RABBIT_ROUTING_KEY" default:"messages.collect.go-template"`
	EventBusRabbitQueueName    string `envconfig:"EVENT_BUS_RABBIT_QUEUE_NAME" default:"events"`
	EventBusRabbitNumConsumers int    `envconfig:"EVENT_BUS_RABBIT_NUM_CONSUMERS" default:"10"`

	// Queue for inbound messages
	InboundRabbitURL              string `envconfig:"INBOUND_RABBIT_URL" default:"amqp://localhost"`
	InboundRabbitExchangeName     string `envconfig:"INBOUND_RABBIT_EXCHANGE_NAME" default:"inbound"`
	InboundRabbitRoutingKey       string `envconfig:"INBOUND_RABBIT_ROUTING_KEY" default:"inbound"`
	InboundRabbitQueueName        string `envconfig:"INBOUND_RABBIT_QUEUE_NAME" default:"inbound"`
	InboundRabbitNumConsumers     int    `envconfig:"INBOUND_RABBIT_NUM_CONSUMERS" default:"10"`
	InboundRabbitQosPrefetchCount int    `envconfig:"INBOUND_RABBIT_QOS_PREFETCH_COUNT" default:"0"`
	InboundRabbitQosPrefetchSize  int    `envconfig:"INBOUND_RABBIT_QOS_PREFETCH_SIZE" default:"0"`
}

func New() *Config {
	return &Config{}
}

func (c *Config) LoadEnvVars() error {
	if err := envconfig.Process(EnvConfigPrefix, c); err != nil {
		return fmt.Errorf("unable to fetch env vars: %s", err)
	}

	return nil
}
