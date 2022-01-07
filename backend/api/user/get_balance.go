package user

import (
	"main/utils"
	"net/http"
)

type GetBalanceReq struct {
	ID string `json:"id"`
}

func (u *user) GetBalance(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w, r)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var req GetBalanceReq
	err := utils.ReadBody(r.Body, &req)
	if err != nil {
		utils.InternalServerError(w, "Read Body", err)
		return
	}

	user, err := u.userOrmer.Select(req.ID)
	if err != nil {
		utils.InternalServerError(w, "Get User", err)
		return
	}

	utils.OK(w, user.Balance)
}
