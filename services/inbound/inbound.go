package inbound

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/proto"
	newrelic "github.com/newrelic/go-agent"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"

	"github.com/batchcorp/go-template/backends/rabbitmq"
	"github.com/batchcorp/go-template/pbevents"
)

const (
	DefaultNumConsumers = 10
)

type IInbound interface {
	ConsumeFunc(msg amqp.Delivery) error
	StartConsumers() error
}

type Inbound struct {
	*Config

	log *logrus.Entry
}

type Config struct {
	Rabbit       rabbitmq.IRabbitMQ
	Context      context.Context
	NRApp        newrelic.Application
	NumConsumers int
}

func New(cfg *Config) (*Inbound, error) {
	if err := validateConfig(cfg); err != nil {
		return nil, fmt.Errorf("unable to validate input cfg: %s", err)
	}

	// Set defaults
	if cfg.Context == nil {
		cfg.Context = context.Background()
	}

	if cfg.NumConsumers == 0 {
		cfg.NumConsumers = DefaultNumConsumers
	}

	return &Inbound{
		Config: cfg,
		log:    logrus.WithField("pkg", "inbound"),
	}, nil
}

func validateConfig(cfg *Config) error {
	if cfg.Rabbit == nil {
		return errors.New("Rabbit cannot be nil")
	}

	return nil
}

func (i *Inbound) StartConsumers() error {
	i.log.Debugf("Launching '%d' inbound consumers", i.NumConsumers)

	for n := 0; n < i.NumConsumers; n++ {
		go i.Rabbit.ConsumeAndRun(i.ConsumeFunc)
	}

	return nil
}

func (i *Inbound) ConsumeFunc(msg amqp.Delivery) error {
	if err := msg.Ack(false); err != nil {
		i.log.Errorf("Error acknowledging message: %s", err)
		return nil
	}

	pbMessage := &pbevents.Intake{}

	if err := proto.Unmarshal(msg.Body, pbMessage); err != nil {
		i.log.Errorf("unable to unmarshal inbound message: %s", err)
		return nil
	}

	_, err := json.Marshal(pbMessage.Headers)
	if err != nil {
		i.log.Errorf("Unable to unmarshal pbMessage.Headers: %s", err)
		return nil
	}

	// Do something with the message

	return nil
}
