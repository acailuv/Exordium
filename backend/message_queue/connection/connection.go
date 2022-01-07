package message_queue

import (
	"main/utils"

	"github.com/streadway/amqp"
)

func NewConnection() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		utils.LogError(utils.GetCurrentCodePosition(), err, nil)
		return nil
	}

	return conn
}
