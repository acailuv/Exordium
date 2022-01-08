package consumer

import (
	"log"
	"main/utils"
)

func (c *consumer) ConsumeUser() {
	q, err := c.channel.QueueDeclare(
		"golang-queue", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		utils.LogError(utils.GetCurrentCodePosition(), err, nil)
		return
	}

	msgs, err := c.channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		utils.LogError(utils.GetCurrentCodePosition(), err, nil)
		return
	}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
}
