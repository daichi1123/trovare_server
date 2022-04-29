package service

import (
	"time"
)

type Session struct {
	ID        int       `json:"id"`
	UUID      string    `json:"uuid"`
	Email     string    `json:"email"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
