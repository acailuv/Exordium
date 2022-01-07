package models

import "time"

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
	Balance   int64     `json:"balance" pg:",use_zero"`
}
