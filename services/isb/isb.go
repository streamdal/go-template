package isb

import (
	"fmt"

	"github.com/batchcorp/schemas/build/go/events"
	"github.com/golang/protobuf/proto"
	newrelic "github.com/newrelic/go-agent"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"

	"github.com/batchcorp/go-template/backends/rabbitmq"
)

const (
	DefaultNumConsumers = 10
)

type IISB interface {
	ConsumeFunc(msg amqp.Delivery) error
	StartConsumers() error
}

type ISB struct {
	*Config

	log *logrus.Entry
}

type Config struct {
	Rabbit       rabbitmq.IRabbitMQ
	NRApp        newrelic.Application
	NumConsumers int
}

func New(cfg *Config) (*ISB, error) {
	if err := validateConfig(cfg); err != nil {
		return nil, fmt.Errorf("unable to validate input cfg: %s", err)
	}

	if cfg.NumConsumers == 0 {
		cfg.NumConsumers = DefaultNumConsumers
	}

	return &ISB{
		Config: cfg,
		log:    logrus.WithField("pkg", "event"),
	}, nil
}

func (i *ISB) StartConsumers() error {
	i.log.Debugf("Launching '%d' event consumers", i.NumConsumers)

	for n := 0; n < i.NumConsumers; n++ {
		go i.Rabbit.ConsumeAndRun(i.ConsumeFunc)
	}

	return nil
}

func validateConfig(cfg *Config) error {
	if cfg.Rabbit == nil {
		return errors.New("Rabbit cannot be nil")
	}

	return nil
}

// This method is intended to be passed as a closure into a rabbit ConsumeAndRun
func (i *ISB) ConsumeFunc(msg amqp.Delivery) error {
	txn := i.NRApp.StartTransaction("ConsumeFunc", nil, nil)
	defer txn.End()

	if err := msg.Ack(false); err != nil {
		i.log.Errorf("Error acknowledging message: %s", err)
		return nil
	}

	pbMessage := &events.Message{}

	if err := proto.Unmarshal(msg.Body, pbMessage); err != nil {
		i.log.Errorf("unable to unmarshal event message: %s", err)
		return nil
	}

	switch pbMessage.Type {
	default:
		i.log.Debug("got an internal message: %+v", pbMessage)
	}

	return nil
}
