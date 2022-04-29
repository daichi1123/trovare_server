package service

import (
	"time"
)

type Owner struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	email     string    `json:"name" `
	password  string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
