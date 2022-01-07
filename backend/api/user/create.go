package user

import (
	"main/database/models"
	"main/utils"
	"net/http"
	"time"
)

type CreateReq struct {
	ID        string `json:"id"`
	CreatedBy string `json:"created_by"`
}

func (u *user) Create(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w, r)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var req CreateReq
	err := utils.ReadBody(r.Body, &req)
	if err != nil {
		utils.InternalServerError(w, "Read Body", err)
		return
	}

	user := models.User{
		ID:        req.ID,
		CreatedAt: time.Now(),
		CreatedBy: req.CreatedBy,
		UpdatedAt: time.Now(),
		UpdatedBy: req.CreatedBy,
		Balance:   0,
	}

	err = u.userOrmer.Insert(user)
	if err != nil {
		utils.InternalServerError(w, "Create", err)
		return
	}

	utils.OK(w, user.ID)
}
