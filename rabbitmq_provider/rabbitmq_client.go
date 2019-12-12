package rabbitmq_provider

import (
	"Ptncafe.Golang.RabbitMQ.Test/constant"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"log"
	"time"
)

type rabbitMqClient struct{
	Connection *amqp.Connection
	Channel *amqp.Channel
}
var _rabbitMqClient *rabbitMqClient

func NewRabbitMqClient(connectionString string, isReuse bool) (*rabbitMqClient ,error){
	if isReuse == true && _rabbitMqClient != nil  {
		return _rabbitMqClient, nil
	}
	rabbitMqConnection, err := amqp.Dial(connectionString)
	if err != nil {
		log.Fatalf("InitConnectionRabbitMq %+v" , errors.Wrap(err,"InitConnectionRabbitMq Dial") )
		return nil,err
	}
	go func() {
		log.Printf("closing: %s", <-rabbitMqConnection.NotifyClose(make(chan *amqp.Error)))
	}()
	rabbitMqChannel, err := rabbitMqConnection.Channel()
	if err != nil {
		log.Fatalf("InitConnectionRabbitMq %+v" , errors.Wrap(err,"InitConnectionRabbitMq Channel") )
		return nil,err
	}
	_rabbitMqClient = &rabbitMqClient{
		rabbitMqConnection,
		rabbitMqChannel,
	}
	return _rabbitMqClient, nil
}

func(rabbitMqClient *rabbitMqClient) Publish(queueName string, data interface{})error{
	dataJson,_:=json.Marshal(data)

	err := rabbitMqClient.Channel.Publish(queueName, "", false,false, amqp.Publishing{
		ContentType: "application/json",
		Body:        dataJson,
		Timestamp:    time.Now(),
	})
	return err
}
func(rabbitMqClient *rabbitMqClient) PublishError(queueName string, data interface{})error{
	dataJson,_:=json.Marshal(data)

	err := rabbitMqClient.Channel.Publish(constant.QueueNameError, "", false,false, amqp.Publishing{
		ContentType: "application/json",
		Body:        dataJson,
		Timestamp:    time.Now(),
		Headers:amqp.Table{"queueName":queueName},
	})
	return err
}

func(rabbitMqClient *rabbitMqClient)  GetChannel(isReuse bool) (*amqp.Channel,error){

	if isReuse {
		return _rabbitMqClient.Channel, nil
	}

	rabbitMqChannel, err := rabbitMqClient.Connection.Channel()
	if err != nil {
		log.Fatalf("InitConnectionRabbitMq %+v" , errors.Wrap(err,"InitConnectionRabbitMq Channel") )
		return nil,err
	}
	//defer rabbitMqChannel.Close()
	return rabbitMqChannel, nil
}

