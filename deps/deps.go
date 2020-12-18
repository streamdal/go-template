package deps

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"time"

	"github.com/InVisionApp/go-health"
	gllogrus "github.com/InVisionApp/go-logger/shims/logrus"
	"github.com/batchcorp/rabbit"
	"github.com/batchcorp/schemas/build/go/events"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"

	"github.com/batchcorp/go-template/backends/cache"
	"github.com/batchcorp/go-template/backends/db"
	"github.com/batchcorp/go-template/backends/etcd"
	"github.com/batchcorp/go-template/backends/kafka"
	"github.com/batchcorp/go-template/backends/postgres"
	"github.com/batchcorp/go-template/config"
	"github.com/batchcorp/go-template/services/hsb"
	"github.com/batchcorp/go-template/services/isb"
)

const (
	DefaultHealthCheckIntervalSecs = 1
)

type customCheck struct{}

type Dependencies struct {
	// Backends
	EtcdBackend         etcd.IEtcd
	ISBDedicatedBackend rabbit.IRabbit
	HSBBackend          kafka.IKafka
	CacheBackend        cache.ICache
	Postgres            *postgres.Postgres

	// Services
	ISBService isb.IISB
	HSBService hsb.IHSB

	HSBChan chan *events.Manifest

	Health         health.IHealth
	DefaultContext context.Context
}

func New(cfg *config.Config) (*Dependencies, error) {
	gohealth := health.New()
	gohealth.Logger = gllogrus.New(nil)

	d := &Dependencies{
		Health:         gohealth,
		DefaultContext: context.Background(),
		HSBChan:        make(chan *events.Manifest, 0),
	}

	if err := d.setupHealthChecks(); err != nil {
		return nil, errors.Wrap(err, "unable to setup health check(s)")
	}

	if err := d.Health.Start(); err != nil {
		return nil, errors.Wrap(err, "unable to start health runner")
	}

	if err := d.setupBackends(cfg); err != nil {
		return nil, errors.Wrap(err, "unable to setup backends")
	}

	if err := d.setupServices(cfg); err != nil {
		return nil, errors.Wrap(err, "unable to setup services")
	}

	return d, nil
}

func (d *Dependencies) setupHealthChecks() error {
	cc := &customCheck{}

	err := d.Health.AddChecks([]*health.Config{
		{
			Name:     "health-check",
			Checker:  cc,
			Interval: time.Duration(DefaultHealthCheckIntervalSecs) * time.Second,
			Fatal:    true,
		},
	})

	if err != nil {
		return err
	}

	return nil
}

func (d *Dependencies) setupBackends(cfg *config.Config) error {
	// CacheBackend k/v store
	cb, err := cache.New()
	if err != nil {
		return errors.Wrap(err, "unable to create new cache instance")
	}

	d.CacheBackend = cb

	// Events rabbitmq backend
	isbDedicatedBackend, err := rabbit.New(&rabbit.Options{
		URL:               cfg.ISBDedicatedURL,
		Mode:              0,
		QueueName:         cfg.ISBDedicatedQueueName,
		ExchangeName:      cfg.ISBDedicatedExchangeName,
		ExchangeType:      amqp.ExchangeTopic,
		ExchangeDeclare:   cfg.ISBDedicatedExchangeDeclare,
		RoutingKey:        cfg.ISBDedicatedRoutingKey,
		RetryReconnectSec: rabbit.DefaultRetryReconnectSec,
		QueueDurable:      cfg.ISBDedicatedQueueDurable,
		QueueExclusive:    cfg.ISBDedicatedQueueExclusive,
		QueueAutoDelete:   cfg.ISBDedicatedQueueAutoDelete,
		QueueDeclare:      cfg.ISBDedicatedQueueDeclare,
		AutoAck:           cfg.ISBDedicatedAutoAck,
		UseTLS:            cfg.ISBDedicatedUseTLS,
		SkipVerifyTLS:     cfg.ISBDedicatedSkipVerifyTLS,
	})
	if err != nil {
		return errors.Wrap(err, "unable to create new rabbit backend")
	}

	d.ISBDedicatedBackend = isbDedicatedBackend

	if cfg.HSBUseTLS {
		logrus.Debug("using TLS for HSB")
	}

	hsbBackend, err := kafka.New(
		&kafka.Options{
			Topic:     cfg.HSBTopicName,
			Brokers:   cfg.HSBBrokerURLs,
			Timeout:   cfg.HSBConnectTimeout,
			BatchSize: cfg.HSBBatchSize,
			UseTLS:    cfg.HSBUseTLS,
		},
		d.DefaultContext,
	)
	if err != nil {
		return errors.Wrap(err, "unable to create new kafka instance")
	}

	d.HSBBackend = hsbBackend

	// etcd
	var tlsConfig *tls.Config

	if cfg.EtcdUseTLS {
		var err error

		tlsConfig, err = createTLSConfig(cfg.EtcdTLSCACert, cfg.EtcdTLSClientCert, cfg.EtcdTLSClientKey)
		if err != nil {
			return errors.Wrap(err, "unable to create TLS config for etcd")
		}
	}

	e, err := etcd.New(&etcd.Options{
		Endpoints:   cfg.EtcdEndpoints,
		DialTimeout: time.Duration(cfg.EtcdDialTimeoutSeconds) * time.Second,
		TLS:         tlsConfig,
	})
	if err != nil {
		return errors.Wrap(err, "unable to create etcd instance")
	}

	d.EtcdBackend = e

	storage, err := db.New(&db.Options{
		Host: cfg.BackendStorageHost,
		Name: cfg.BackendStorageName,
		User: cfg.BackendStorageUser,
		Pass: cfg.BackendStoragePass,
		Port: cfg.BackendStoragePort,
	})

	if err != nil {
		return errors.Wrap(err, "unable to create new db instance")
	}

	pg := postgres.New(storage)

	d.Postgres = pg

	return nil
}

func (d *Dependencies) setupServices(cfg *config.Config) error {
	isbService, err := isb.New(&isb.Config{
		RabbitMap: map[string]*isb.RabbitConfig{
			"dedicated": {
				RabbitInstance: d.ISBDedicatedBackend,
				NumConsumers:   cfg.ISBDedicatedNumConsumers,
				Func:           "DedicatedConsumeFunc",
			},
		},
	})
	if err != nil {
		return errors.Wrap(err, "unable to setup event")
	}

	d.ISBService = isbService

	hsbService, err := hsb.New(&hsb.Config{
		Kafka:         d.HSBBackend,
		NumPublishers: cfg.HSBNumPublishers,
		HSBChan:       d.HSBChan,
	})
	if err != nil {
		return errors.Wrap(err, "unable to setup hsb")
	}

	d.HSBService = hsbService

	return nil
}

func createTLSConfig(caCert, clientCert, clientKey string) (*tls.Config, error) {
	cert, err := tls.X509KeyPair([]byte(clientCert), []byte(clientKey))
	if err != nil {
		return nil, errors.Wrap(err, "unable to load cert + key")
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM([]byte(caCert))

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}, nil
}

// Satisfy the go-health.ICheckable interface
func (c *customCheck) Status() (interface{}, error) {
	if false {
		return nil, errors.New("something major just broke")
	}

	// You can return additional information pertaining to the check as long
	// as it can be JSON marshalled
	return map[string]int{}, nil
}
