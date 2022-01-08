package user

import (
	ormer "main/database/ormer/user"
	"main/rabbitmq/publisher"
	"net/http"

	"github.com/go-pg/pg/v10"
	"github.com/go-redis/redis"
	"github.com/streadway/amqp"
)

type User interface {
	Create(w http.ResponseWriter, r *http.Request)
	AddBalance(w http.ResponseWriter, r *http.Request)
	Withdraw(w http.ResponseWriter, r *http.Request)
	GetBalance(w http.ResponseWriter, r *http.Request)
	PublishUser(w http.ResponseWriter, r *http.Request)
	SetRedis(w http.ResponseWriter, r *http.Request)
	GetRedis(w http.ResponseWriter, r *http.Request)
}

type user struct {
	userOrmer ormer.UserOrmer
	publisher publisher.Publisher
	redis     redis.Client
}

func NewUserClient(db *pg.DB, rabbitMQ *amqp.Connection, redis *redis.Client) User {
	return &user{
		userOrmer: ormer.NewUserOrmer(db),
		publisher: publisher.NewPublisherClient(rabbitMQ),
		redis:     *redis,
	}
}
