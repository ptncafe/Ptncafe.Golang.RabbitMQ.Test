package consumer

import (
	"Ptncafe.Golang.RabbitMQ.Test/constant"
	"Ptncafe.Golang.RabbitMQ.Test/model"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"log"
	"time"
)

func ShopConsumer(connection string){
	handler := HandlerConsumer(func(data *[]byte) error {
		shopData := model.StoreDto{}
		_ = json.Unmarshal(*data, &shopData)
		log.Printf("ShopConsumer %v",spew.Sdump(shopData))
		time.Sleep(time.Duration(1000) * time.Millisecond)
		return nil
	})
	CommonConsumer(connection,constant.QueueNameShop,1000, handler)
}
