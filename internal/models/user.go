package models

import "time"

type User struct {
	UserID        int64     `json:"user_id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	Password_hash string    `json:"-"`
	Created_at    time.Time `json:"created_at"`
}
