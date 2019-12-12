package consumer

import (
	"Ptncafe.Golang.RabbitMQ.Test/constant"
	"Ptncafe.Golang.RabbitMQ.Test/model"
	"Ptncafe.Golang.RabbitMQ.Test/rabbitmq_provider"
	"encoding/json"


	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	"log"
	"time"
)

func ShopConsumer(connection string) error{
	handler := HandlerConsumer(func(data *[]byte) error {
		shopData := model.StoreDto{}
		_ = json.Unmarshal(*data, &shopData)
		log.Printf("ShopConsumer %v",spew.Sdump(shopData))
		time.Sleep(100 * time.Millisecond)
		return nil
	})
	clientRabbitMq,err:= rabbitmq_provider.NewRabbitMqClient(constant.RabbitMqConnectionString,false)
	if err!= nil {
		log.Print(errors.Wrap(err,"ShopConsumer"))
		return nil
	}
	_, err = NewConsumer(clientRabbitMq,constant.QueueNameShop,5, handler)
	if err != nil {
		return errors.Wrap(err, "ShopConsumer")
	}
	return nil
}
