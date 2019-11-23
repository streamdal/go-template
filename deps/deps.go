package deps

import (
	"context"
	"fmt"
	"time"

	"github.com/InVisionApp/go-health"
	gllogrus "github.com/InVisionApp/go-logger/shims/logrus"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/batchcorp/go-template/backends/rabbitmq"
	"github.com/batchcorp/go-template/backends/workers"
	"github.com/batchcorp/go-template/config"
)

const (
	DefaultHealthCheckIntervalSecs = 1
	PoolSize                       = 10
)

var (
	log *logrus.Entry
)

type customCheck struct{}

type Dependencies struct {
	Health         health.IHealth
	RabbitMQ       rabbitmq.IRabbitMQ
	DefaultContext context.Context
}

func init() {
	log = logrus.WithField("pkg", "deps")
}

func New(cfg *config.Config) (*Dependencies, error) {
	gohealth := health.New()
	gohealth.Logger = gllogrus.New(nil)

	d := &Dependencies{
		Health:         gohealth,
		DefaultContext: context.Background(),
	}

	if err := d.setupHealthChecks(); err != nil {
		return nil, err
	}

	if err := d.Health.Start(); err != nil {
		return nil, err
	}

	if err := d.setupBackends(cfg); err != nil {
		return nil, err
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
	r, err := rabbitmq.New(cfg, d.DefaultContext)

	if err != nil {
		return err
	}

	d.RabbitMQ = r
	return nil
}

func (d *Dependencies) RunWorkers() {
	for i := 0; i < PoolSize; i++ {
		id := uuid.New()
		worker := workers.NewWorker(id, d.RabbitMQ.GetConsumerChannel(), d.DefaultContext)
		go worker.Run()
	}
}

// Satisfy the go-health.ICheckable interface
func (c *customCheck) Status() (interface{}, error) {
	if false {
		return nil, fmt.Errorf("Something major just broke.")
	}

	// You can return additional information pertaining to the check as long
	// as it can be JSON marshalled
	return map[string]int{}, nil
}
