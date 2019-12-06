package consumer

import (
	"Ptncafe.Golang.RabbitMQ.Test/rabbitmq_provider"
	"github.com/pkg/errors"
	"log"
	"os"
)
type HandlerConsumer func(*[]byte) error

func CommonConsumer(connection string, queueName string, prefetchCount int,handler HandlerConsumer){
	channel,err := rabbitmq_provider.GetRabbitMqChannel(connection)
	if err != nil {
		log.Fatalf("ShopConsumer %s %+v ", connection, errors.Wrap(err,"ShopConsumer") )
	}
	defer channel.Close()

	err = channel.Qos(prefetchCount, 0, false)

	messageChannel, err := channel.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	stopChan := make(chan bool)

	go func() {
		log.Printf("Consumer ready, PID: %d", os.Getpid())
		for d := range messageChannel {
			err := handler(&d.Body)
			if err == nil {
				err = d.Ack(false)
				if err !=nil {
					log.Printf("CommonConsumer %v", errors.Wrap(err,"CommonConsumer"))
				}
			}
		}
	}()

	// Stop for program termination
	<-stopChan

}