package websocket

import (
	"main/utils"
	"net/http"

	gorillaWebsocket "github.com/gorilla/websocket"
)

func (ws *websocket) WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := gorillaWebsocket.Upgrader{}
	upgrader.CheckOrigin = utils.RequestOriginUpgraderOverride

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.LogError(utils.GetCurrentCodePosition(), err, nil)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			utils.LogError(utils.GetCurrentCodePosition(), err, nil)
			return
		}

		utils.LogInfo(utils.GetCurrentCodePosition(), "Message Type:", messageType, "Message:", message)
		err = conn.WriteMessage(messageType, []byte("Pong!"))
		if err != nil {
			utils.LogError(utils.GetCurrentCodePosition(), err, nil)
			return
		}
	}
}
