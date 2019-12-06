package main

import (
	"Ptncafe.Golang.RabbitMQ.Test/constant"
	"Ptncafe.Golang.RabbitMQ.Test/controller"
	"Ptncafe.Golang.RabbitMQ.Test/rabbitmq_provider"
	"github.com/gin-gonic/gin"
)

func main() {
	channelRabbitMq, err := rabbitmq_provider.InitConnectionRabbitMq(constant.RabbitMqConnectionString)
	///defer channelRabbitMq.Close()
	if err !=nil {
		panic(err)
	}
	err = rabbitmq_provider.InitRabbitMq(channelRabbitMq)
	if err !=nil {
		panic(err)
	}
	
	r := gin.Default()
	r.GET("queue", controller.QueueController)
	r.Run()
}
