package main

import (
	"fmt"
	"log"
	"net/http"

	"main/api/healthcheck"
	"main/api/user"
	database "main/database/connection"
	message_queue "main/message_queue/connection"
	"main/message_queue/consumer"

	"github.com/gorilla/mux"
)

type handlerClients struct {
	user        user.User
	healthcheck healthcheck.Healthcheck
}

func handleRequests(clients handlerClients) {
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/create", clients.user.Create).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/add-balance", clients.user.AddBalance).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/withdraw", clients.user.Withdraw).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/get-balance", clients.user.GetBalance).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/publish-user", clients.user.PublishUser).Methods(http.MethodPost, http.MethodOptions)

	router.HandleFunc("/healthcheck", clients.healthcheck.StatusCheck).Methods(http.MethodGet, http.MethodOptions)

	log.Fatal(http.ListenAndServe(":5000", router))
}

func main() {

	db := database.NewConnection()
	rabbitMQ := message_queue.NewConnection()

	userClient := user.NewUserClient(db, rabbitMQ)
	healthcheckClient := healthcheck.NewHealthcheckClient()

	clients := handlerClients{
		user:        userClient,
		healthcheck: healthcheckClient,
	}

	consumer := consumer.NewConsumerClient(rabbitMQ)
	consumer.RegisterAllConsumer()

	fmt.Println("Listening at *:5000")
	handleRequests(clients)

}
