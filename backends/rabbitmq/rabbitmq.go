// Generic rabbitmq backend used for both internal event consumption and hsb
// message publishing.
package rabbitmq

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"github.com/relistan/go-director"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

//go:generate counterfeiter -o ../../fakes/rabbitmq/rabbitmq.go . IRabbitMQ

type IRabbitMQ interface {
	ConsumeAndRun(runFunc func(msg amqp.Delivery) error)
	Publish([]byte) error
}

type RabbitMQ struct {
	Client          *amqp.Connection
	RabbitMQChannel *amqp.Channel
	Queue           amqp.Queue
	WorkerChannel   chan string
	DefaultContext  context.Context
	Looper          *director.FreeLooper

	log *logrus.Entry
}

type Options struct {
	URL              string
	QueueName        string // A blank queue name makes sense in _some_ cases (such as for fanout exchange)
	ExchangeName     string
	ExchangeType     string
	RoutingKey       string
	QosPrefetchCount int
	QosPrefetchSize  int
}

func New(opts *Options, ctx context.Context) (*RabbitMQ, error) {
	if err := validateOptions(opts); err != nil {
		return nil, errors.Wrap(err, "invalid options")
	}

	ac, err := amqp.Dial(opts.URL)
	if err != nil {
		return nil, err
	}

	ch, err := ac.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "Channel instantiation failure")
	}

	if err := ch.Qos(opts.QosPrefetchCount, opts.QosPrefetchSize, false); err != nil {
		return nil, errors.Wrap(err, "unable to set qos policy")
	}

	if err := ch.ExchangeDeclare(
		opts.ExchangeName,
		opts.ExchangeType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return nil, errors.Wrap(err, "unable to declare exchange")
	}

	eventsQueue, err := ch.QueueDeclare(
		opts.QueueName,
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		return nil, err
	}

	if err := ch.QueueBind(
		opts.QueueName,
		opts.RoutingKey,
		opts.ExchangeName,
		false,
		nil,
	); err != nil {
		return nil, errors.Wrap(err, "unable to bind queue")
	}

	return &RabbitMQ{
		log:             logrus.WithField("pkg", "backends.rabbitmq"),
		Client:          ac,
		RabbitMQChannel: ch,
		Queue:           eventsQueue,
		WorkerChannel:   make(chan string),
		DefaultContext:  ctx,
		Looper:          director.NewFreeLooper(director.FOREVER, make(chan error)),
	}, nil
}

func validateOptions(opts *Options) error {
	if opts.URL == "" {
		return errors.New("URL cannot be empty")
	}

	if opts.ExchangeType == "" {
		return errors.New("ExchangeType cannot be empty")
	}

	if opts.ExchangeName == "" {
		return errors.New("ExchangeName cannot be empty")
	}

	if opts.RoutingKey == "" {
		return errors.New("RoutingKey cannot be empty")
	}

	return nil
}

// Consume events from queue and run given func
func (r *RabbitMQ) ConsumeAndRun(runFunc func(msg amqp.Delivery) error) {
	messageChannel, err := r.RabbitMQChannel.Consume(
		r.Queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		r.log.Fatal("unable to start consuming messages from %s queue: %s",
			r.Queue.Name, err)
	}

	r.log.Debugf("Consumer ready, PID: %d", os.Getpid())

	r.Looper.Loop(func() error {
		select {
		case msg := <-messageChannel:
			if err := runFunc(msg); err != nil {
				return err
			}
		case <-r.DefaultContext.Done():
			r.log.Warning("received notice to quit loop")
			r.Looper.Quit()
		}
		return nil
	})

	r.log.Debug("exiting")
}

func (r *RabbitMQ) Publish(body []byte) error {
	if err := r.RabbitMQChannel.Publish("", r.Queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		// TODO: Do we need to specify ContentType?
		Body: body,
	}); err != nil {
		return err
	}

	r.log.Debugf("Publishing body %s", string(body))

	return nil
}
