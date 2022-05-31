package models

import "time"

type (
	Drink struct {
		ID           int       `json:"id,omitempty"`
		Name         string    `json:"name"`
		Description  string    `json:"description"`
		ImageURL     string    `json:"image_url"`
		RestaurantID int       `json:"restaurant_id"`
		CreatedAt    time.Time `json:"-"`
		UpdatedAt    time.Time `json:"-"`
		DeletedAt    time.Time `json:"-"`
	}
)
