package controller

import (
	"Ptncafe.Golang.RabbitMQ.Test/constant"
	"Ptncafe.Golang.RabbitMQ.Test/model"
	"Ptncafe.Golang.RabbitMQ.Test/rabbitmq_provider"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"log"
	"time"
)

func QueueController(ctx *gin.Context){
	var storeDto = model.StoreDto{
		Id:1,
		Name: "name",
		Code: "code",
		ShopStatus: 2,
		UpdatedDate: time.Now(),
		}
	clientRabbitMq,err:= rabbitmq_provider.NewRabbitMqClient(constant.RabbitMqConnectionString, true)
	if err!= nil {
		log.Print(errors.Wrap(err,"QueueController"))
		ctx.JSON(500, gin.H{
			"error": fmt.Sprintf("QueueController %v", err),
		})
		return
	}
	err = clientRabbitMq.Publish(constant.QueueNameShop , storeDto)
	if err!= nil {
		log.Print(errors.Wrap(err,"QueueController"))
		ctx.JSON(500, gin.H{
			"error": fmt.Sprintf("QueueController %v", err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}