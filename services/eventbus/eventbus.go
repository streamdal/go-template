package eventbus

import (
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

type IEventBus interface {
	ConsumeFunc(msg amqp.Delivery) error
	StartConsumers() error
}

type EventBus struct {
	*Config

	log *logrus.Entry
}

type Config struct {
	Rabbit       rabbitmq.IRabbitMQ
	NRApp        newrelic.Application
	NumConsumers int
}

func New(cfg *Config) (*EventBus, error) {
	if err := validateConfig(cfg); err != nil {
		return nil, fmt.Errorf("unable to validate input cfg: %s", err)
	}

	if cfg.NumConsumers == 0 {
		cfg.NumConsumers = DefaultNumConsumers
	}

	return &EventBus{
		Config: cfg,
		log:    logrus.WithField("pkg", "event"),
	}, nil
}

func (e *EventBus) StartConsumers() error {
	e.log.Debugf("Launching '%d' event consumers", e.NumConsumers)

	for n := 0; n < e.NumConsumers; n++ {
		go e.Rabbit.ConsumeAndRun(e.ConsumeFunc)
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
func (e *EventBus) ConsumeFunc(msg amqp.Delivery) error {
	txn := e.NRApp.StartTransaction("ConsumeFunc", nil, nil)
	defer txn.End()

	if err := msg.Ack(false); err != nil {
		e.log.Errorf("Error acknowledging message: %s", err)
		return nil
	}

	pbMessage := &pbevents.Message{}

	if err := proto.Unmarshal(msg.Body, pbMessage); err != nil {
		e.log.Errorf("unable to unmarshal event message: %s", err)
		return nil
	}

	switch pbMessage.Type {
	default:
		e.log.Debug("got an internal message: %+v", pbMessage)
	}

	return nil
}
