package service

import "time"

type RestaurantGenre struct {
	ID           int       `json:"id"`
	RestaurantID int       `json:"restaurant_id"`
	GenreID      int       `json:"genre_id"`
	CreatedAt    time.Time `json:"created_at"`
}
