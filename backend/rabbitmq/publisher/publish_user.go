package publisher

import (
	"main/database/models"
	"main/utils"

	"github.com/streadway/amqp"
)

func (p *publisher) PublishUser(user models.User) error {
	q, err := p.channel.QueueDeclare(
		"golang-queue", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		utils.LogError(utils.GetCurrentCodePosition(), err, nil)
		return err
	}

	body := utils.ConvertToByteArray(user)
	err = p.channel.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		utils.LogError(utils.GetCurrentCodePosition(), err, nil)
		return err
	}

	utils.LogInfo(utils.GetCurrentCodePosition(), "Message Sent!", user)
	return nil
}
