package user

import (
	"fmt"
	"main/database/models"
	"main/utils"
	"net/http"
	"time"
)

type WithdrawReq struct {
	ID        string `json:"id"`
	UpdatedBy string `json:"updated_by"`
	Amount    int64  `json:"amount"`
}

func (u *user) Withdraw(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w, r)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var req WithdrawReq
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

	if userData.Balance < req.Amount {
		utils.BadRequest(w, "Insufficient Balance!")
		return
	}

	newUserData := models.User{
		ID:        req.ID,
		UpdatedAt: time.Now(),
		UpdatedBy: req.UpdatedBy,
		Balance:   userData.Balance - req.Amount,
	}

	err = u.userOrmer.Update(newUserData)
	if err != nil {
		utils.InternalServerError(w, "Withdraw", err)
		return
	}

	if req.Amount == 1 {
		utils.OK(w, fmt.Sprintf("%d Coin has been withdrawn successfully!", req.Amount))
	} else {
		utils.OK(w, fmt.Sprintf("%d Coins have been withdrawn successfully!", req.Amount))
	}
}
