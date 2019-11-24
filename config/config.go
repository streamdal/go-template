package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

const (
	EnvConfigPrefix = "GO_TEMPLATE"
)

type Config struct {
	ListenAddress string `envconfig:"LISTEN_ADDRESS" default:":8080"`
	RabbitMQURL   string `envconfig:"RABBITMQ_URL" required:"true"`
	HealthFreqSec int    `envconfig:"HEALTH_FREQ_SEC" default:"60"`
	EnvName       string `envconfig:"ENV_NAME" default:"dev"`
	ServiceName   string `envconfig:"SERVICE_NAME" default:"go-template"`

	EventsQueueName string `envconfig:"EVENTS_QUEUE_NAME" default:"events"`
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
