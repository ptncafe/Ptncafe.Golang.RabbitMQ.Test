package consumer

import (
	"Ptncafe.Golang.RabbitMQ.Test/rabbitmq_provider"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"log"
)
type HandlerConsumer func(*[]byte) error


type Consumer struct {
	queueName string
}

func NewConsumer(amqpURI string, queueName string, refetchCount int,handler HandlerConsumer)  error {
	rabbitMqConnection, err := amqp.Dial(amqpURI)
	if err != nil {
		log.Fatalf("InitConnectionRabbitMq %+v" , errors.Wrap(err,"InitConnectionRabbitMq Dial") )
		return err
	}
	go func() {
		log.Printf("closing: %s", <-rabbitMqConnection.NotifyClose(make(chan *amqp.Error)))
	}()
	for i := 0; i < refetchCount; i++ {
		done :=    make(chan error)
		channel, err := rabbitMqConnection.Channel()
		if err != nil {
			return fmt.Errorf("Channel: %s", err)
		}
		err = channel.Qos(1, 0, false)
		if err != nil {
			return fmt.Errorf("Qos: %s", err)
		}
		log.Printf("Queue bound to Exchange, starting Consume (consumer tag %q)", queueName)

		deliveries, err := channel.Consume(
			queueName, // name
			queueName,      // consumerTag,
			false,      // noAck
			false,      // exclusive
			false,      // noLocal
			false,      // noWait
			nil,        // arguments
		)
		if err != nil {
			return fmt.Errorf("Queue Consume: %s", err)
		}

		go handle(deliveries, done,queueName,handler)
	}
	return nil
}

func handle(deliveries <-chan amqp.Delivery, done chan error,queueName string,handler HandlerConsumer) {
	log.Printf("handle total %v %s", len(deliveries), spew.Sdump(deliveries))
	for d := range deliveries {
			err:= handler(&d.Body)
			if err != nil {
				go func() {
					rabbitmq_provider.PublishError(queueName, d.Body)
				}()
			} else{
				go func() {
					d.Ack(false)
				}()

			}
	}

	//d:= <-deliveries
	//log.Printf("handle %v", string(d.Body))
	//d.Ack(false)
	//close(done)

	log.Printf("handle: deliveries channel closed")
	done <- nil
}