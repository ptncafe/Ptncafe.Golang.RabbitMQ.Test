package main

import (
	"Ptncafe.Golang.RabbitMQ.Test/constant"
	"Ptncafe.Golang.RabbitMQ.Test/consumer"
	"Ptncafe.Golang.RabbitMQ.Test/controller"
	"Ptncafe.Golang.RabbitMQ.Test/rabbitmq_provider"
	"github.com/gin-gonic/gin"
)

var rabbitMqClient rabbitmq_provider.IRabbitMqClient
func main() {

	err := rabbitmq_provider.RegisterQueue()
	if err !=nil {
		panic(err)
	}
	go func() {
		consumer.InitConsumer(constant.RabbitMqConnectionString)
	}()

	r := gin.Default()
	r.GET("queue", controller.QueueController)
	r.Run()
}
