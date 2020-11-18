package config

import (
	"errors"
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
	ServiceName            string   `envconfig:"SERVICE_NAME" default:"go-template"`
	EtcdEndpoints          []string `envconfig:"ETCD_ENDPOINTS"default:"localhost:2379"`
	EtcdDialTimeoutSeconds int      `envconfig:"ETCD_DIAL_TIMEOUT_SECONDS" default:"5"`
	EtcdUseTLS             bool     `envconfig:"ETCD_USE_TLS" default:"false"`
	EtcdTLSCACert          string   `envconfig:"ETCD_TLS_CA_CERT"`
	EtcdTLSClientCert      string   `envconfig:"ETCD_TLS_CLIENT_CERT"`
	EtcdTLSClientKey       string   `envconfig:"ETCD_TLS_CLIENT_KEY"`

	// Queue for _internal_ events
	ISBDedicatedURL               string `envconfig:"ISB_DEDICATED_URL" default:"amqp://localhost"`
	ISBDedicatedExchangeName      string `envconfig:"ISB_DEDICATED_EXCHANGE_NAME" default:"events"`
	ISBDedicatedExchangeDeclare   bool   `envconfig:"ISB_DEDICATED_EXCHANGE_DECLARE" default:"true"`
	ISBDedicatedExchangeDurable   bool   `envconfig:"ISB_DEDICATED_EXCHANGE_DURABLE" default:"true"`
	ISBDedicatedRoutingKey        string `envconfig:"ISB_DEDICATED_ROUTING_KEY" default:"messages.collect.#"` // UPDATE THIS
	ISBDedicatedQueueName         string `envconfig:"ISB_DEDICATED_QUEUE_NAME" default:""`                    // UPDATE THIS
	ISBDedicatedNumConsumers      int    `envconfig:"ISB_DEDICATED_NUM_CONSUMERS" default:"10"`
	ISBDedicatedRetryReconnectSec int    `envconfig:"ISB_DEDICATED_RETRY_RECONNECT_SEC" default:"10"`
	ISBDedicatedAutoAck           bool   `envconfig:"ISB_DEDICATED_AUTO_ACK" default:"false"`
	ISBDedicatedQueueDeclare      bool   `envconfig:"ISB_DEDICATED_QUEUE_DECLARE" default:"true"`
	ISBDedicatedQueueDurable      bool   `envconfig:"ISB_DEDICATED_QUEUE_DURABLE" default:"false"`
	ISBDedicatedQueueExclusive    bool   `envconfig:"ISB_DEDICATED_QUEUE_EXCLUSIVE" default:"true"`
	ISBDedicatedQueueAutoDelete   bool   `envconfig:"ISB_DEDICATED_QUEUE_AUTO_DELETE" default:"true"`

	// Queue for hsb messages
	HSBBrokerURLs     []string      `envconfig:"HSB_BROKER_URLS" default:"localhost:9092"`
	HSBUseTLS         bool          `envconfig:"HSB_USE_TLS" default:"false"`
	HSBTopicName      string        `envconfig:"HSB_TOPIC_NAME" default:"inbound"`
	HSBNumPublishers  int           `envconfig:"HSB_NUM_PUBLISHERS" default:"10"`
	HSBConnectTimeout time.Duration `envconfig:"HSB_CONNECT_TIMEOUT" default:"10s"`
	HSBBatchSize      int           `envconfig:"HSB_BATCH_SIZE" default:"1"`

	// Medium term storage
	BackendStorageHost string `envconfig:"BACKEND_STORAGE_HOST" default:"localhost"`
	BackendStorageName string `envconfig:"BACKEND_STORAGE_NAME" default:"go-template"`
	BackendStoragePort int    `envconfig:"BACKEND_STORAGE_PORT" default:"5432"`
	BackendStorageUser string `envconfig:"BACKEND_STORAGE_USER" default:"postgres"`
	BackendStoragePass string `envconfig:"BACKEND_STORAGE_PASS"`
}

func New() *Config {
	return &Config{}
}

func (c *Config) LoadEnvVars() error {
	if err := envconfig.Process(EnvConfigPrefix, c); err != nil {
		return fmt.Errorf("unable to fetch env vars: %s", err)
	}

	if c.EtcdUseTLS {
		if c.EtcdTLSCACert == "" {
			return errors.New("ETCD_TLS_CA_CERT must be set when ETCD_USE_TLS set to true")
		}

		if c.EtcdTLSClientCert == "" {
			return errors.New("ETCD_TLS_CLIENT_CERT must be set when ETCD_USE_TLS set to true")
		}

		if c.EtcdTLSClientKey == "" {
			return errors.New("ETCD_TLS_CLIENT_KEY must be set when ETCD_USE_TLS set to true")
		}
	}

	return nil
}
