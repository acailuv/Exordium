package user

import (
	ormer "main/database/ormer/user"
	"main/message_queue/publisher"
	"net/http"

	"github.com/go-pg/pg/v10"
	"github.com/streadway/amqp"
)

type User interface {
	Create(w http.ResponseWriter, r *http.Request)
	AddBalance(w http.ResponseWriter, r *http.Request)
	Withdraw(w http.ResponseWriter, r *http.Request)
	GetBalance(w http.ResponseWriter, r *http.Request)
	PublishUser(w http.ResponseWriter, r *http.Request)
}

type user struct {
	userOrmer ormer.UserOrmer
	publisher publisher.Publisher
}

func NewUserClient(db *pg.DB, rabbitMQ *amqp.Connection) User {
	return &user{
		userOrmer: ormer.NewUserOrmer(db),
		publisher: publisher.NewPublisherClient(rabbitMQ),
	}
}
