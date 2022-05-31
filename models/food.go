package models

import "time"

type (
	Food struct {
		id           int       `json:"id,omitempty"`
		Name         string    `json:"name"`
		Description  string    `json:"description"`
		ImageURL     string    `json:"image_url"`
		RestaurantID int       `json:"restaurant_id"`
		CreatedAt    time.Time `json:"-"`
		UpdatedAt    time.Time `json:"-"`
		DeletedAt    time.Time `json:"-"`
	}
)
