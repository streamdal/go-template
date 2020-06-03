package deps

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/InVisionApp/go-health"
	gllogrus "github.com/InVisionApp/go-logger/shims/logrus"
	"github.com/batchcorp/schemas/build/go/events"
	newrelic "github.com/newrelic/go-agent"
	"github.com/newrelic/go-agent/_integrations/nrlogrus"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"

	"github.com/batchcorp/go-template/backends/kafka"
	"github.com/batchcorp/go-template/backends/rabbitmq"
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
	ISBBackend rabbitmq.IRabbitMQ
	HSBBackend kafka.IKafka

	// Services
	ISBService isb.IISB
	HSBService hsb.IHSB

	HSBChan chan *events.Manifest

	Health         health.IHealth
	DefaultContext context.Context
	NRApp          newrelic.Application
}

func New(cfg *config.Config) (*Dependencies, error) {
	gohealth := health.New()
	gohealth.Logger = gllogrus.New(nil)

	d := &Dependencies{
		Health:         gohealth,
		DefaultContext: context.Background(),
		HSBChan:        make(chan *events.Manifest, 0),
	}

	if err := d.setupNR(cfg); err != nil {
		return nil, errors.Wrap(err, "unable to setup New Relic agent")
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
	// Events rabbitmq backend
	eventsRabbit, err := rabbitmq.New(
		&rabbitmq.Options{
			URL:          cfg.ISBURL,
			ExchangeType: amqp.ExchangeTopic,
			ExchangeName: cfg.ISBExchangeName,
			RoutingKey:   cfg.ISBRoutingKey,
			QueueName:    cfg.ISBQueueName,
		},
		d.DefaultContext,
	)
	if err != nil {
		return errors.Wrap(err, "unable to create new events rabbit instance")
	}

	d.ISBBackend = eventsRabbit

	hsb, err := kafka.New(
		&kafka.Options{
			Topic:     cfg.HSBTopicName,
			Brokers:   cfg.HSBBrokerURLs,
			Timeout:   cfg.HSBConnectTimeout,
			BatchSize: cfg.HSBBatchSize,
		},
		d.DefaultContext,
	)
	if err != nil {
		return errors.Wrap(err, "unable to create new kafka instance")
	}

	d.HSBBackend = hsb

	return nil
}

func (d *Dependencies) setupServices(cfg *config.Config) error {
	isb, err := isb.New(&isb.Config{
		Rabbit:       d.ISBBackend,
		NRApp:        d.NRApp,
		NumConsumers: cfg.ISBNumConsumers,
	})
	if err != nil {
		return errors.Wrap(err, "unable to setup event")
	}

	d.ISBService = isb

	hsb, err := hsb.New(&hsb.Config{
		Kafka:         d.HSBBackend,
		NRApp:         d.NRApp,
		NumPublishers: cfg.HSBNumPublishers,
		HSBChan:       d.HSBChan,
	})
	if err != nil {
		return errors.Wrap(err, "unable to setup hsb")
	}

	d.HSBService = hsb

	return nil
}

func (d *Dependencies) setupNR(cfg *config.Config) error {
	appName := fmt.Sprintf("%s-%s", cfg.ServiceName, cfg.EnvName)

	nrConfig := newrelic.NewConfig(appName, cfg.NewRelicLicense)

	if cfg.NewRelicLicense == "" {
		logrus.Info("no license provided; disabling new relic")
		nrConfig.Enabled = false
	}

	nrConfig.ErrorCollector.IgnoreStatusCodes = []int{
		http.StatusNotFound,
	}

	// Force NR agent to log at INFO level (as its debug logs are very loud)
	l := logrus.New()
	l.SetLevel(logrus.InfoLevel)
	nrConfig.Logger = nrlogrus.Transform(l)

	app, err := newrelic.NewApplication(nrConfig)
	if err != nil {
		return err
	}

	d.NRApp = app

	return nil
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
