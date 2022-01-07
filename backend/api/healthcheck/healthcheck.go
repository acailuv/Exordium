package healthcheck

import (
	"net/http"
)

type Healthcheck interface {
	StatusCheck(w http.ResponseWriter, r *http.Request)
}

type healthcheck struct {
}

func NewHealthcheckClient() Healthcheck {
	return &healthcheck{}
}
