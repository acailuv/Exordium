package healthcheck

import (
	"main/utils"
	"net/http"
)

func (h *healthcheck) StatusCheck(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w, r)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	utils.OK(w, "All System Operational.")
}
