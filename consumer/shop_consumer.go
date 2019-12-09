package consumer

import (
	"Ptncafe.Golang.RabbitMQ.Test/constant"
	"Ptncafe.Golang.RabbitMQ.Test/model"
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
		time.Sleep(200 * time.Millisecond)
		return nil
	})

	_,err := NewConsumer(connection,constant.QueueNameShop,5, handler)
	if err != nil {
		return errors.Wrap(err, "ShopConsumer")
	}
	return nil
}
