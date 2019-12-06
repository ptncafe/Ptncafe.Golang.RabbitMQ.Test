package controller

import (
	"Ptncafe.Golang.RabbitMQ.Test/constant"
	"Ptncafe.Golang.RabbitMQ.Test/model"
	"Ptncafe.Golang.RabbitMQ.Test/rabbitmq_provider"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func QueueController(ctx *gin.Context){
	var channel,err = rabbitmq_provider.RabbitMqChannel()
	//defer channel.Close()
	if err!= nil {
		log.Print(errors.Wrap(err,"QueueController"))
		ctx.JSON(500, gin.H{
			"error": fmt.Sprintf("QueueController %v", err),
		})
		return
	}
	var storeDto = model.StoreDto{
		Id:1,
		Name: "name",
		Code: "code",
		ShopStatus: 2,
		}
	dataJson,_:=json.Marshal(storeDto)
	err = channel.Publish(constant.QueueNameShop, "", false,false, amqp.Publishing{
		ContentType: "application/json",
		Body:        dataJson,
		Timestamp:    time.Now(),
	})
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