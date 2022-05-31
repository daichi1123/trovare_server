package trovare

import (
	"go_api/configs"
	"go_api/handler/genre"
	"go_api/handler/location"
	"go_api/handler/restaurant"
	"net/http"

	"github.com/rs/cors"
)

func Router() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	})
	handler := cors.Default().Handler(mux)

	// search
	mux.HandleFunc("/v1/restaurants/search", restaurant.SearchRestaurants)

	// routing for restaurants
	//restaurant.Restaurant.CreateTestRestaurant(restaurant.Restaurant{})
	mux.HandleFunc("/v1/restaurants", restaurant.GetAllRestaurants)
	mux.HandleFunc("/v1/restaurants/", restaurant.GetRestaurant)
	mux.HandleFunc("/v1/restaurants/create", restaurant.CreateRestaurant)
	mux.HandleFunc("/v1/restaurants/update/", restaurant.UpdateRestaurantInfo)
	mux.HandleFunc("/v1/restaurants/delete/", restaurant.DeleteRestaurant)

	mux.HandleFunc("/v1/restaurants/get/list", restaurant.GetRestaurantsForList)

	// routing for locations
	mux.HandleFunc("/v1/locations/post/current", location.PostCurrentLocation)

	// routing of genre
	mux.HandleFunc("/v1/genres/create", genre.CreateGenre)

	// routing for user
	http.ListenAndServe(configs.Config.Port, handler)
	return nil
}
