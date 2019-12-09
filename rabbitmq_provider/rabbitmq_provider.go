package rabbitmq_provider

import (
	"Ptncafe.Golang.RabbitMQ.Test/constant"
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
	go func() {
		log.Printf("closing: %s", <-rabbitMqConnection.NotifyClose(make(chan *amqp.Error)))
	}()
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
	go func() {
		log.Printf("closing: %s", <-rabbitMqConnection.NotifyClose(make(chan *amqp.Error)))
	}()
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
func PublishError(queueName string, data interface{})error{
	dataJson,_:=json.Marshal(data)

	err := rabbitMqChannel.Publish(constant.QueueNameError, "", false,false, amqp.Publishing{
		ContentType: "application/json",
		Body:        dataJson,
		Timestamp:    time.Now(),
		Headers:amqp.Table{"queueName":queueName},
	})
	return err
}
