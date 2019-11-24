package rabbitmq

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/relistan/go-director"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"

	"github.com/batchcorp/go-template/config"
)

const (
	BatchSize = 1
)

//go:generate counterfeiter -o ../../fakes/rabbitmq/rabbitmq.go . IRabbitMQ

type IRabbitMQ interface {
	Get() error
	Listen()
	Publish([]byte)
	GetConsumerChannel() chan string
}

type RabbitMQ struct {
	log             *logrus.Entry
	Client          *amqp.Connection
	RabbitMQChannel *amqp.Channel
	Queue           amqp.Queue
	WorkerChannel   chan string
	prefetchCount   int
	prefetchSize    int
	DefaultContext  context.Context
	Looper          *director.FreeLooper
}

func New(cfg *config.Config, ctx context.Context) (*RabbitMQ, error) {
	ac, err := amqp.Dial(cfg.RabbitMQURL)
	if err != nil {
		return nil, err
	}

	ch, err := ac.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "Channel instantiation failure")
	}

	queue, err := ch.QueueDeclare("events", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		log:             logrus.WithField("pkg", "backends.rabbitmq"),
		Client:          ac,
		RabbitMQChannel: ch,
		Queue:           queue,
		WorkerChannel:   make(chan string),
		DefaultContext:  ctx,
		Looper:          director.NewFreeLooper(director.FOREVER, make(chan error)),
	}, nil
}

func (r *RabbitMQ) Get() error {
	return nil
}

func (r *RabbitMQ) Listen() {
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
		fmt.Println(err)
	}

	r.log.Printf("Consumer ready, PID: %d", os.Getpid())

	r.Looper.Loop(func() error {
		select {
		case msg := <-messageChannel:
			r.WorkerChannel <- string(msg.Body)
			if err := msg.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message")
			}
		case <-r.DefaultContext.Done():
			r.Looper.Quit()
		}
		return nil
	})

}

func (r *RabbitMQ) Publish(body []byte) {
	err := r.RabbitMQChannel.Publish("", r.Queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	if err != nil {
		r.log.Fatalf("Error publishing message: %s", err.Error())
	}

	r.log.Printf("Publishing body %s", string(body))
}

func (r *RabbitMQ) GetConsumerChannel() chan string {
	return r.WorkerChannel
}
