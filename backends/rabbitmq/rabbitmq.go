// Generic rabbitmq backend used for both internal event consumption and hsb
// message publishing.
package rabbitmq

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/relistan/go-director"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

//go:generate counterfeiter -o ../../fakes/rabbitmq/rabbitmq.go . IRabbitMQ

const (
	DefaultRetryReconnectSec = 60
)

type IRabbitMQ interface {
	ConsumeAndRun(runFunc func(msg amqp.Delivery) error)
	Publish([]byte) error
}

type RabbitMQ struct {
	Conn                    *amqp.Connection
	ConsumerServerChannel   *amqp.Channel
	ConsumerDeliveryChannel <-chan amqp.Delivery
	ConsumerRWMutex         *sync.RWMutex // Controls access for both RabbitMQ*Channels maps
	NotifyCloseChan         chan *amqp.Error
	ProducerServerChannel   *amqp.Channel
	ProducerRWMutex         *sync.RWMutex
	DefaultContext          context.Context
	Looper                  *director.FreeLooper
	Options                 *Options

	log *logrus.Entry
}

type Options struct {
	URL               string
	QueueName         string // A blank queue name makes sense in _some_ cases (such as for fanout exchange)
	ExchangeName      string
	ExchangeType      string
	RoutingKey        string
	QosPrefetchCount  int
	QosPrefetchSize   int
	RetryReconnectSec int
	QueueDurable      bool
	QueueExclusive    bool
	QueueAutoDelete   bool
}

func New(opts *Options, ctx context.Context) (*RabbitMQ, error) {
	if err := validateOptions(opts); err != nil {
		return nil, errors.Wrap(err, "invalid options")
	}

	ac, err := amqp.Dial(opts.URL)
	if err != nil {
		return nil, err
	}

	r := &RabbitMQ{
		log:             logrus.WithField("pkg", "backends.rabbitmq"),
		Conn:            ac,
		ConsumerRWMutex: &sync.RWMutex{},
		NotifyCloseChan: make(chan *amqp.Error),
		ProducerRWMutex: &sync.RWMutex{},
		DefaultContext:  ctx,
		Looper:          director.NewFreeLooper(director.FOREVER, make(chan error)),
		Options:         opts,
	}

	if err := r.newConsumerChannels(); err != nil {
		return nil, errors.Wrap(err, "unable to get initial delivery channel")
	}

	ac.NotifyClose(r.NotifyCloseChan)

	// Launch connection watcher/reconnect
	go r.watchNotifyClose()

	return r, nil
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

	if opts.RetryReconnectSec == 0 {
		opts.RetryReconnectSec = DefaultRetryReconnectSec
	}

	return nil
}

// Consume events from queue and run given func
func (r *RabbitMQ) ConsumeAndRun(runFunc func(msg amqp.Delivery) error) {
	r.Looper.Loop(func() error {
		select {
		case msg := <-r.delivery():
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
	// Is this the first time we're publishing?
	if r.ProducerServerChannel == nil {
		ch, err := r.newServerChannel()
		if err != nil {
			return errors.Wrap(err, "unable to create server channel")
		}

		r.ProducerRWMutex.Lock()
		r.ProducerServerChannel = ch
		r.ProducerRWMutex.Unlock()
	}

	r.ProducerRWMutex.RLock()
	defer r.ProducerRWMutex.RUnlock()

	if err := r.ProducerServerChannel.Publish("", r.Options.QueueName, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		// TODO: Do we need to specify ContentType?
		Body: body,
	}); err != nil {
		return err
	}

	r.log.Debugf("Publishing body %s", string(body))

	return nil
}

func (r *RabbitMQ) watchNotifyClose() {
	for {
		closeErr := <-r.NotifyCloseChan

		r.log.Debugf("received message on notify close channel: '%+v' (reconnecting)", closeErr)

		// Acquire mutex to pause all consumers while we reconnect AND prevent
		// access to the channel map
		r.ConsumerRWMutex.Lock()

		var attempts int

		for {
			attempts++

			if err := r.reconnect(); err != nil {
				r.log.Warningf("unable to complete reconnect: %s; retrying in %d", err, r.Options.RetryReconnectSec)
				time.Sleep(time.Duration(r.Options.RetryReconnectSec) * time.Second)
				continue
			}

			r.log.Debugf("successfully reconnected after %d attempts", attempts)
			break
		}

		// Create and set a new notify close channel (since old one gets closed)
		r.NotifyCloseChan = make(chan *amqp.Error, 0)
		r.Conn.NotifyClose(r.NotifyCloseChan)

		// Update channel
		if err := r.newConsumerChannels(); err != nil {
			logrus.Errorf("unable to set new channel: %s", err)

			// TODO: 1. This is super shitty 2. Need to push an error to NR about this.
			panic(fmt.Sprintf("unable to set new channel: %s", err))
		}

		// Unlock so that consumers can begin reading messages from a new channel
		r.ConsumerRWMutex.Unlock()

		r.log.Debug("watchNotifyClose has completed successfully")
	}
}

func (r *RabbitMQ) newServerChannel() (*amqp.Channel, error) {
	if r.Conn == nil {
		return nil, errors.New("r.Conn is nil - did this get instantiated correctly? bug?")
	}

	ch, err := r.Conn.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "unable to instantiate channel")
	}

	if err := ch.Qos(r.Options.QosPrefetchCount, r.Options.QosPrefetchSize, false); err != nil {
		return nil, errors.Wrap(err, "unable to set qos policy")
	}

	if err := ch.ExchangeDeclare(
		r.Options.ExchangeName,
		r.Options.ExchangeType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return nil, errors.Wrap(err, "unable to declare exchange")
	}

	if _, err := ch.QueueDeclare(
		r.Options.QueueName,
		r.Options.QueueDurable,
		r.Options.QueueAutoDelete,
		r.Options.QueueExclusive,
		false,
		nil,
	); err != nil {
		return nil, err
	}

	if err := ch.QueueBind(
		r.Options.QueueName,
		r.Options.RoutingKey,
		r.Options.ExchangeName,
		false,
		nil,
	); err != nil {
		return nil, errors.Wrap(err, "unable to bind queue")
	}

	return ch, nil
}

func (r *RabbitMQ) newConsumerChannels() error {
	serverChannel, err := r.newServerChannel()
	if err != nil {
		return errors.Wrap(err, "unable to create new server channel")
	}

	deliveryChannel, err := serverChannel.Consume(r.Options.QueueName,
		"",
		false,
		r.Options.QueueExclusive,
		false,
		false,
		nil,
	)

	if err != nil {
		return errors.Wrap(err, "unable to create delivery channel")
	}

	r.ConsumerServerChannel = serverChannel
	r.ConsumerDeliveryChannel = deliveryChannel

	return nil
}

func (r *RabbitMQ) reconnect() error {
	ac, err := amqp.Dial(r.Options.URL)
	if err != nil {
		return err
	}

	r.Conn = ac

	return nil
}

func (r *RabbitMQ) delivery() <-chan amqp.Delivery {
	// Acquire lock (in case we are reconnecting and channels are being swapped)
	r.ConsumerRWMutex.RLock()
	defer r.ConsumerRWMutex.RUnlock()

	return r.ConsumerDeliveryChannel
}
