package user

import (
	"main/database/models"
	"main/utils"
	"net/http"
)

func (u *user) PublishUser(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w, r)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var req models.User
	err := utils.ReadBody(r.Body, &req)
	if err != nil {
		utils.InternalServerError(w, "Read Body", err)
		return
	}

	err = u.publisher.PublishUser(req)
	if err != nil {
		utils.InternalServerError(w, "Publish User", err)
		return
	}

	utils.OK(w, "User Message Published Successfully!")
}
