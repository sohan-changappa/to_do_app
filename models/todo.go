package models

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Title     string    `json:title`
	Completed bool      `json:completed`
	CreatedAt time.Time `json:created_at`
}
