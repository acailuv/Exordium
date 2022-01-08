package user

import (
	"fmt"
	"main/utils"
	"net/http"
)

type GetRedisReq struct {
	Key string `json:"key"`
}

func (u *user) GetRedis(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w, r)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var req GetRedisReq
	err := utils.ReadBody(r.Body, &req)
	if err != nil {
		utils.InternalServerError(w, "Read Body", err)
		return
	}

	result, err := u.redis.Get(req.Key).Result()
	if err != nil {
		utils.InternalServerError(w, "Redis Get", err)
		return
	}

	utils.OK(w, fmt.Sprintf("Data Get! Key: %v; Value: %v", req.Key, result))
}
