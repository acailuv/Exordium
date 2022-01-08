package consumer

import (
	"main/utils"

	"github.com/streadway/amqp"
)

type Consumer interface {
	ConsumeUser()
	RegisterAllConsumer()
}

type consumer struct {
	rabbitMQ *amqp.Connection
	channel  *amqp.Channel
}

func NewConsumerClient(rabbitMQ *amqp.Connection) Consumer {
	channel, err := rabbitMQ.Channel()
	if err != nil {
		utils.LogError(utils.GetCurrentCodePosition(), err, nil)
	}

	return &consumer{
		rabbitMQ: rabbitMQ,
		channel:  channel,
	}
}
