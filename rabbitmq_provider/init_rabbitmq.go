package rabbitmq_provider

import (
	"Ptncafe.Golang.RabbitMQ.Test/constant"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"log"
)

var queueNames = []string{
	constant.QueueNameError,
	constant.QueueNameShop,
	constant.QueueShopConfigShippingFee,
}

// InitRabbitMq mind set là exchange và queue sẽ dùng tên, 1 exchange và 1 queue
func RegisterQueue() error{
	rabbitMqClient, err := NewRabbitMqClient(constant.RabbitMqConnectionString, true)
	if err !=nil {
		panic(err)
	}
	rabbitMqChannel,err :=rabbitMqClient.GetChannel(true)
	if err !=nil {
		panic(err)
	}
	for _, name := range  queueNames{
		_ = InitExchange(rabbitMqChannel, name)
		_ = InitQueue(rabbitMqChannel,name)
		//return err
	}
	return nil
}

func InitExchange(channel *amqp.Channel, name string) error{
	err := channel.ExchangeDeclare(name, "fanout",true,false, false, false, nil)
	if err != nil {
		log.Fatalf("InitExchange %s %+v ", name, errors.Wrap(err,"InitExchange") )
		return err
	}
	return nil
}


func InitQueue(channel *amqp.Channel, name string) error{
	_, err := channel.QueueDeclare(name,true,false, false, false, nil)
	if err != nil {
		log.Fatalf("InitQueue %s %+v ", name, errors.Wrap(err,"InitQueue") )
		return err
	}
	err = channel.QueueBind(
		name, // queue name
		"",     // routing key
		name, // exchange
		false,
		nil)
	if err != nil {
		log.Fatalf("InitQueue %s %+v ", name, errors.Wrap(err,"InitQueue") )
		return err
	}
	log.Printf("InitQueue %s done ", name )

	return nil
}