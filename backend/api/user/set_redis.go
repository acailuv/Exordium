package user

import (
	"fmt"
	"main/utils"
	"net/http"
)

type SetRedisReq struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (u *user) SetRedis(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w, r)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var req SetRedisReq
	err := utils.ReadBody(r.Body, &req)
	if err != nil {
		utils.InternalServerError(w, "Read Body", err)
		return
	}

	err = u.redis.Set(req.Key, req.Value, 0).Err()
	if err != nil {
		utils.InternalServerError(w, "Redis Set", err)
		return
	}

	utils.OK(w, fmt.Sprintf("Data inserted successfully! Key: %v; Value: %v", req.Key, req.Value))
}
