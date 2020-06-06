package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

const (
	EnvConfigPrefix = "GO_TEMPLATE"
)

type Config struct {
	ListenAddress          string   `envconfig:"LISTEN_ADDRESS" default:":8282"`
	HealthFreqSec          int      `envconfig:"HEALTH_FREQ_SEC" default:"60"`
	EnvName                string   `envconfig:"ENV_NAME" default:"local"`
	BadgerDirectory        string   `envconfig:"BADGER_DIRECTORY" default:"./backend-data/badger"`
	ServiceName            string   `envconfig:"SERVICE_NAME" default:"go-template"`
	EtcdEndpoints          []string `envconfig:"ETCD_ENDPOINTS"default:"localhost:2379"`
	EtcdDialTimeoutSeconds int      `envconfig:"ETCD_DIAL_TIMEOUT_SECONDS" default:"5"`
	EtcdTLSEnabled         bool     `envconfig:"ETCD_TLS_ENABLED" default:"false"`
	EtcdTLSCACert          string   `envconfig:"ETCD_TLS_CA_CERT"`
	EtcdTLSClientCert      string   `envconfig:"ETCD_TLS_CLIENT_CERT"`
	EtcdTLSClientKey       string   `envconfig:"ETCD_TLS_CLIENT_KEY"`
	NewRelicLicense        string   `envconfig:"NEW_RELIC_LICENSE"`

	// Queue for _internal_ events
	ISBURL          string `envconfig:"ISB_URL" default:"amqp://localhost"`
	ISBExchangeName string `envconfig:"ISB_EXCHANGE_NAME" default:"events"`
	ISBRoutingKey   string `envconfig:"ISB_ROUTING_KEY" default:"messages.collect.go-template"`
	ISBQueueName    string `envconfig:"ISB_QUEUE_NAME" default:"events"`
	ISBNumConsumers int    `envconfig:"ISB_NUM_CONSUMERS" default:"10"`

	// Queue for hsb messages
	HSBBrokerURLs     []string      `envconfig:"HSB_BROKER_URLS" default:"localhost:9092"`
	HSBUseTLS         bool          `envconfig:"HSB_USE_TLS" default:"false"`
	HSBTopicName      string        `envconfig:"HSB_TOPIC_NAME" default:"hsb"`
	HSBNumPublishers  int           `envconfig:"HSB_NUM_PUBLISHERS" default:"10"`
	HSBConnectTimeout time.Duration `envconfig:"HSB_CONNECT_TIMEOUT" default:"10s"`
	HSBBatchSize      int           `envconfig:"HSB_BATCH_SIZE" default:"1"`
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
