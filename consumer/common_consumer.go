package consumer

import (
	"Ptncafe.Golang.RabbitMQ.Test/rabbitmq_provider"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/streadway/amqp"
	"log"
)
type HandlerConsumer func(*[]byte) error

type ConmonConsumer struct {
	rabbitMqClient rabbitmq_provider.IRabbitMqClient
	queueName string
}
//Qos koh chạy, phải gắn 1 connecion, nhiều channel
func NewConsumer(rabbitMqClient rabbitmq_provider.IRabbitMqClient, queueName string, refetchCount int,handler HandlerConsumer)  (*ConmonConsumer,error) {
	var conmonConsumer = ConmonConsumer{
		rabbitMqClient: rabbitMqClient,
		queueName:queueName,
	}

	for i := 0; i < refetchCount; i++ {
		done :=    make(chan error)
		channel,err := rabbitMqClient.GetChannel(false)
		//defer channel.Close()
		if err != nil {
			return nil,fmt.Errorf("Channel: %s", err)
		}
		err = channel.Qos(1, 0, false)
		if err != nil {
			return nil, fmt.Errorf("Qos: %s", err)
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
			return nil, fmt.Errorf("Queue Consume: %s", err)
		}

		go handle(conmonConsumer, deliveries, done,queueName,handler)
	}
	return &conmonConsumer, nil
}

func handle(conmonConsumer ConmonConsumer,deliveries <-chan amqp.Delivery, done chan error,queueName string,handler HandlerConsumer) {
	log.Printf("handle total %v %s", len(deliveries), spew.Sdump(deliveries))
	for d := range deliveries {
			err:= handler(&d.Body)
			//err:= errors.New("Lỗi test")
			if err != nil {
				log.Printf("ConmonConsumer %v", err)
			 	err = conmonConsumer.rabbitMqClient.PublishError(queueName, d.Body)
				if err != nil {
					log.Printf("ConmonConsumer PublishError %v", err)
				}else{
					d.Ack(false)
				}
			}else{
				d.Ack(false)
			}
	}

	//d:= <-deliveries
	//log.Printf("handle %v", string(d.Body))
	//d.Ack(false)
	//close(done)

	log.Printf("handle: deliveries channel closed")
	done <- nil
}

