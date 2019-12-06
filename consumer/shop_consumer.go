package consumer

import (
	"Ptncafe.Golang.RabbitMQ.Test/constant"
	"Ptncafe.Golang.RabbitMQ.Test/model"
	"Ptncafe.Golang.RabbitMQ.Test/rabbitmq_provider"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	"log"
	"os"
	"time"
)

func ShopConsumer(connection string){
	channel,err := rabbitmq_provider.GetRabbitMqChannel(connection)
	if err != nil {
		log.Fatalf("ShopConsumer %s %+v ", connection, errors.Wrap(err,"ShopConsumer") )
	}
	defer channel.Close()

	err = channel.Qos(1, 0, false)

	messageChannel, err := channel.Consume(
		constant.QueueNameShop,
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
			log.Printf("Received a message: %s", d.Body)

			dataJson := &model.StoreDto{}

			err := json.Unmarshal(d.Body, dataJson)
			time.Sleep(time.Duration(5) * time.Second)
			if err != nil {
				err = rabbitmq_provider.Publish(constant.QueueNameError, dataJson)
				if err !=nil {
					log.Printf("Error Publish : %s", err)
				}
			}

			log.Printf("Result of %s",spew.Sdump( dataJson))

			if err := d.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Consumer message ack done %s",spew.Sdump( dataJson))
			}

		}
	}()

	// Stop for program termination
	<-stopChan

}
