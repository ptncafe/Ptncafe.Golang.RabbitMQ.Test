package rabbitmq_provider

import (
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"log"
)

var rabbitMqConnection *amqp.Connection
var rabbitMqChannel *amqp.Channel
func InitConnectionRabbitMq(connectionString string) (*amqp.Channel ,error){
	var err error
	rabbitMqConnection, err = amqp.Dial(connectionString)
	if err != nil {
		log.Fatalf("InitConnectionRabbitMq %+v" , errors.Wrap(err,"InitConnectionRabbitMq Dial") )
		return nil,err
	}
	rabbitMqChannel, err = rabbitMqConnection.Channel()
	if err != nil {
		log.Fatalf("InitConnectionRabbitMq %+v" , errors.Wrap(err,"InitConnectionRabbitMq Channel") )
		return nil,err
	}
	return rabbitMqChannel,nil
}

func RabbitMqChannel() (*amqp.Channel,error){
	return rabbitMqChannel,nil
}
