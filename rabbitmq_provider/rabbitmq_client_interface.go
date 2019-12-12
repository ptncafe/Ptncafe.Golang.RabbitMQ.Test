package rabbitmq_provider

import "github.com/streadway/amqp"

type IRabbitMqClient interface {
	Publish(queueName string, data interface{})error
	PublishError(queueName string, data interface{})error
	GetChannel(isReuse bool) (*amqp.Channel, error)
}
