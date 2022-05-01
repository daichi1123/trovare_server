package service

import "time"

type (
	Genre struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"-"`
		UpdatedAt time.Time `json:"-"`
	}

	CreateGenreRequest struct {
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"-"`
	}

	CreateGenreResponse struct {
		Genre Genre `json:"genre"`
	}
)
