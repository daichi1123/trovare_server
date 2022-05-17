package router

import (
	"go_api/cmd/pkg"
	"go_api/configs"
	"go_api/internal/trovare/auth"
	"go_api/internal/trovare/handler/genre"
	"go_api/internal/trovare/handler/restaurant"
	"go_api/internal/trovare/handler/user"
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
	mux.HandleFunc("/v1/signin", auth.Signin)

	//下記RedirectIndexはエラーが出る
	//mux.HandleFunc("/", utils.RedirectIndex)

	// routing for restaurants
	restaurant.Restaurant.CreateTestRestaurant(restaurant.Restaurant{})
	mux.HandleFunc("/v1/restaurants", restaurant.GetAllRestaurants)
	mux.HandleFunc("/v1/restaurants/", restaurant.GetRestaurant)
	mux.HandleFunc("/v1/restaurants/create", restaurant.CreateRestaurant)
	mux.HandleFunc("/v1/restaurants/update/", restaurant.UpdateRestaurantInfo)
	mux.HandleFunc("/v1/restaurants/delete/", restaurant.DeleteRestaurant)

	mux.HandleFunc("/v1/restaurants/get/list", restaurant.GetRestaurantsForList)

	// routing of genre
	mux.HandleFunc("/v1/genres/create", genre.CreateGenre)

	// routing for user
	mux.HandleFunc("/index", pkg.Index)
	mux.HandleFunc("/v1/signup", auth.Signup)
	mux.HandleFunc("/v1/signin/test", user.Signin)
	mux.HandleFunc("/v1/login", user.Login)
	mux.HandleFunc("/v1/check-session", user.Authentication)

	mux.HandleFunc("/after-login", auth.AfterLogin)

	http.ListenAndServe(configs.Config.Port, handler)
	return nil
}
