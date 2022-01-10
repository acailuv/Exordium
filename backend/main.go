package main

import (
	"fmt"
	"log"
	"net/http"

	"main/api/healthcheck"
	"main/api/user"
	"main/api/websocket"
	database "main/database/connection"
	rabbitmq "main/rabbitmq/connection"
	"main/rabbitmq/consumer"
	redis "main/redis/connection"

	"github.com/gorilla/mux"
)

type handlerClients struct {
	user        user.User
	healthcheck healthcheck.Healthcheck
	websocket   websocket.Websocket
}

func handleRequests(clients handlerClients) {
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/create", clients.user.Create).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/add-balance", clients.user.AddBalance).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/withdraw", clients.user.Withdraw).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/get-balance", clients.user.GetBalance).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/publish-user", clients.user.PublishUser).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/set-cache", clients.user.SetRedis).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/get-cache", clients.user.GetRedis).Methods(http.MethodPost, http.MethodOptions)

	router.HandleFunc("/healthcheck", clients.healthcheck.StatusCheck).Methods(http.MethodGet, http.MethodOptions)

	router.HandleFunc("/websocket", clients.websocket.WebsocketHandler)

	log.Fatal(http.ListenAndServe(":5000", router))
}

func main() {

	db := database.NewConnection()
	rabbitMQ := rabbitmq.NewConnection()
	redis := redis.NewConnection()

	userClient := user.NewUserClient(db, rabbitMQ, redis)
	healthcheckClient := healthcheck.NewHealthcheckClient()
	websocketClient := websocket.NewWebsocketClient()

	clients := handlerClients{
		user:        userClient,
		healthcheck: healthcheckClient,
		websocket:   websocketClient,
	}

	consumer := consumer.NewConsumerClient(rabbitMQ)
	consumer.RegisterAllConsumer()

	fmt.Println("Listening at *:5000")
	handleRequests(clients)

}
