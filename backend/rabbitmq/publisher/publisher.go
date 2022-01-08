package publisher

import (
	"main/database/models"
	"main/utils"

	"github.com/streadway/amqp"
)

type Publisher interface {
	PublishUser(models.User) error
}

type publisher struct {
	rabbitMQ *amqp.Connection
	channel  *amqp.Channel
}

func NewPublisherClient(rabbitMQ *amqp.Connection) Publisher {
	channel, err := rabbitMQ.Channel()
	if err != nil {
		utils.LogError(utils.GetCurrentCodePosition(), err, nil)
	}

	return &publisher{
		rabbitMQ: rabbitMQ,
		channel:  channel,
	}
}
