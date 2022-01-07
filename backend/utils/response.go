package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ServerResponse struct {
	IsSuccess bool        `json:"is_success"`
	Status    int         `json:"status"`
	Error     error       `json:"error"`
	Data      interface{} `json:"data"`
}

func InternalServerError(w http.ResponseWriter, process string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ServerResponse{
		IsSuccess: false,
		Status:    http.StatusInternalServerError,
		Error:     err,
		Data:      fmt.Sprintf("Internal Server Error: %v: %v", process, err),
	})
}

func BadRequest(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(ServerResponse{
		IsSuccess: false,
		Status:    http.StatusBadRequest,
		Data:      message,
	})
}

func OK(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ServerResponse{
		IsSuccess: true,
		Status:    http.StatusOK,
		Data:      data,
	})
}
