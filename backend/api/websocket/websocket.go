package websocket

import (
	"net/http"
)

type Websocket interface {
	WebsocketHandler(w http.ResponseWriter, r *http.Request)
}

type websocket struct {
}

func NewWebsocketClient() Websocket {
	return &websocket{}
}
