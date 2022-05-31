package models

import "time"

type (
	Restaurant struct {
		ID          int       `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Rating      int       `json:"rating"`
		ZipCode     int       `json:"zip_code"`
		Address     string    `json:"address"`
		Lat         float64   `json:"lat"`
		Lng         float64   `json:"lng"`
		ImageURL    string    `json:"image_url"`
		OwnerID     int       `json:"owner_id"`
		GenreID     int       `json:"genre_id"`
		CreatedAt   time.Time `json:"-"`
		UpdatedAt   time.Time `json:"-"`
		DeletedAt   time.Time `json:"-"`
	}

	CreateRestaurantsRequest struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		// TODO: testが終了次第ratingを削除するdefault値を決めておく
		Rating    int       `json:"rating"`
		ZipCode   int       `json:"zip_code"`
		Address   string    `json:"address"`
		Lat       float64   `json:"lat"`
		Lng       float64   `json:"lng"`
		ImageURL  string    `json:"image_url"`
		CreatedAt time.Time `json:"-"`
	}

	CreateRestaurantsResponse struct {
		ID          int       `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Rating      int       `json:"rating"`
		ZipCode     int       `json:"zip_code"`
		Address     string    `json:"address"`
		Lat         float64   `json:"lat"`
		Lng         float64   `json:"lng"`
		ImageURL    string    `json:"image_url"`
		CreatedAt   time.Time `json:"-"`
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

	GetRestaurantLocationRequest struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		ZipCode  int    `json:"zip_code"`
		Address  string `json:"address"`
		ImageURL string `json:"image_url"`
	}

	GetRestaurantLocationResponse struct {
		GetRestaurantLocationRequest []GetRestaurantLocationRequest `json:"get_rst_location"`
	}

	CreateRestaurantResponse struct {
		Restaurant Restaurant `json:"restaurant"`
	}

	UpdateRestaurantRequest struct {
	}

	UpdateRestaurantResponse struct {
	}
)
