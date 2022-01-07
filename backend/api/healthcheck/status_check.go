package healthcheck

import (
	"main/utils"
	"net/http"
)

func (h *healthcheck) StatusCheck(w http.ResponseWriter, r *http.Request) {
	utils.OK(w, "All System Operational.")
}
