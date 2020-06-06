package kafka

import (
	"context"
	"crypto/tls"
	"fmt"
	"math/rand"
	"time"

	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

const (
	DefaultConnectTimeout = 10 * time.Second
	DefaultBatchSize      = 10
)

type IKafka interface {
	Publish(key, value []byte) error
}

type Kafka struct {
	Writer  *kafka.Writer
	Options *Options
	Context context.Context
}

type Options struct {
	Topic     string
	Brokers   []string
	Timeout   time.Duration
	BatchSize int
	UseTLS    bool
}

func New(opts *Options, ctx context.Context) (*Kafka, error) {
	if err := validateOptions(opts); err != nil {
		return nil, errors.Wrap(err, "invalid options")
	}

	if ctx == nil {
		ctx = context.Background()
	}

	dialer := &kafka.Dialer{
		Timeout: opts.Timeout,
	}

	if opts.UseTLS {
		dialer.TLS = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	if _, err := dialer.DialContext(ctx, "tcp", opts.Brokers[0]); err != nil {
		return nil, fmt.Errorf("unable to create initial connection to broken '%s': %s",
			opts.Brokers[0], err)
	}

	return &Kafka{
		Writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers:   opts.Brokers,
			Topic:     opts.Topic,
			Dialer:    dialer,
			BatchSize: opts.BatchSize,
		}),
		Options: opts,
		Context: ctx,
	}, nil
}

func validateOptions(opts *Options) error {
	if len(opts.Brokers) == 0 {
		return errors.New("brokers cannot be empty")
	}

	if opts.Topic == "" {
		return errors.New("topic cannot be empty")
	}

	if opts.Timeout == 0 {
		opts.Timeout = DefaultConnectTimeout
	}

	if opts.BatchSize <= 0 {
		opts.BatchSize = DefaultBatchSize
	}

	return nil
}

func (k *Kafka) Publish(key, value []byte) error {
	randKey := rand.Intn(10000)

	logrus.Infof("writing a message to kafka: %d", randKey)

	if err := k.Writer.WriteMessages(k.Context, kafka.Message{
		Key:   key,
		Value: value,
	}); err != nil {
		return errors.Wrap(err, "unable to publish message(s)")
	}

	logrus.Infof("finished writing message to kafka: %d", randKey)

	logrus.Info(k.Writer.Stats())

	return nil
}
