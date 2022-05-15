package restaurant

import (
	"go_api/internal/trovare/service"
)

type (
	User                    service.User
	Session                 service.Session
	Restaurant              service.Restaurant
	GetRestaurantsResponse  service.GetRestaurantsResponse
	GetRestaurantLists      service.GetRestaurantsListsResponse
	CreateRestaurantRequest service.CreateRestaurantRequest
)
