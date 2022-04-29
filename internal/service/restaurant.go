package service

import "time"

type Restaurant struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	RestaurantId int       `json:"restaurant_id"`
	OwnerId      int       `json:"owner_id"`
	Rating       int       `json:"rating"`
	CreatedAt    time.Time `json:"created_at"`
}

type CreateRestaurantRequest struct {
}

type CreateRestaurantResponse struct {
	Restaurant Restaurant `json:"restaurant"`
}

type UpdateRestaurantRequest struct {
}

type UpdateRestaurantResponse struct {
}
