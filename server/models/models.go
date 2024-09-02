package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	Username  string    `json:"user_name"`
	Password  string    `json:"password"` // Omitting the password from JSON output for security
	CreatedAt time.Time `json:"created_at"`
}

type GetUserResp struct {
    UserID    int       `json:"user_id"`
	Username  string    `json:"user_name"`
	Password  string    `json:"password"` // Omitting the password from JSON output for security
	CreatedAt time.Time `json:"created_at"`
}

type GetById struct {
	Id int
}

type Task struct {
	TodoID      int       `json:"todo_id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"` // Omits empty description in JSON
	Completed   bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
}
