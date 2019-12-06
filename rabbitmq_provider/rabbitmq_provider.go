package rabbitmq_provider

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"log"
	"time"
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

func GetRabbitMqChannel(connectionString string)(*amqp.Channel,error){
	rabbitMqConnection, err := amqp.Dial(connectionString)
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

func Publish(queueName string, data interface{})error{
	dataJson,_:=json.Marshal(data)

	err := rabbitMqChannel.Publish(queueName, "", false,false, amqp.Publishing{
		ContentType: "application/json",
		Body:        dataJson,
		Timestamp:    time.Now(),
	})
	return err
}
