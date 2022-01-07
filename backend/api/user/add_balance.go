package user

import (
	"fmt"
	"main/database/models"
	"main/utils"
	"net/http"
	"time"
)

type AddBalanceReq struct {
	ID        string `json:"id"`
	UpdatedBy string `json:"updated_by"`
	Amount    int64  `json:"amount"`
}

func (u *user) AddBalance(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w, r)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var req AddBalanceReq
	err := utils.ReadBody(r.Body, &req)
	if err != nil {
		utils.InternalServerError(w, "Read Body", err)
		return
	}

	userData, err := u.userOrmer.Select(req.ID)
	if err != nil {
		utils.InternalServerError(w, "Select", err)
		return
	}

	newUserData := models.User{
		ID:        userData.ID,
		UpdatedAt: time.Now(),
		UpdatedBy: req.UpdatedBy,
		Balance:   userData.Balance + req.Amount,
	}

	err = u.userOrmer.Update(newUserData)
	if err != nil {
		utils.InternalServerError(w, "Add Balance", err)
		return
	}

	if req.Amount == 1 {
		utils.OK(w, fmt.Sprintf("%d Coin has been added successfully!", req.Amount))
	} else {
		utils.OK(w, fmt.Sprintf("%d Coins have been added successfully!", req.Amount))
	}
}
