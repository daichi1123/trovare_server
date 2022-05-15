package service

import "time"

type (
	Restaurant struct {
		ID          int       `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Rating      int       `json:"rating"`
		ZipCode     int       `json:"zip_code"`
		Address     string    `json:"address"`
		ImageURL    string    `json:"image_url"`
		CreatedAt   time.Time `json:"-"`
		UpdatedAt   time.Time `json:"-"`
		DeletedAt   time.Time `json:"-"`
	}

	GetRestaurantsResponse struct {
		Restaurants []Restaurant `json:"restaurants"`
	}

	GetRestaurantsListsResponse struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		ZipCode     int    `json:"zip_code"`
		Address     string `json:"address"`
		Rating      int    `json:"rating"`
		ImageURL    string `json:"image_url"`
	}

	CreateRestaurantRequest struct {
		Name        string    `json:"name"`
		Description string    `json:"description"`
		ZipCode     int       `json:"zip_code"`
		Address     string    `json:"address"`
		ImageURL    string    `json:"image_url"`
		CreatedAt   time.Time `json:"-"`
	}

	CreateRestaurantResponse struct {
		Restaurant Restaurant `json:"restaurant"`
	}

	UpdateRestaurantRequest struct {
	}

	UpdateRestaurantResponse struct {
	}
)
