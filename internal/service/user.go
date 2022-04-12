package service

import "time"

type User struct {
	ID        int       `json:"id"`
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
	CreatedAt time.Time
}
