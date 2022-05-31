package restaurant

import "go_api/models"

type (
	User                         models.User
	Session                      models.Session
	Restaurant                   models.Restaurant
	GetRestaurantsResponse       models.GetRestaurantsResponse
	GetRestaurantLists           models.GetRestaurantsListsResponse
	GetRestaurantLocationRequest models.GetRestaurantLocationRequest
	CreateRestaurantRequest      models.CreateRestaurantRequest
)
