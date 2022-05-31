package models

import "time"

type (
	Owner struct {
		id        int       `json:"id,omitempty"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		CreatedAt time.Time `json:"-"`
		UpdatedAt time.Time `json:"-"`
		DeletedAt time.Time `json:"-"`
	}
)
