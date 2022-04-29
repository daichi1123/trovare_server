package restaurantHandler

import (
	"encoding/json"
	"go_api/utils"
	"log"
	"net/http"
	"time"
)

// name of method need `_'table name' or not`??

func CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	var restaurant Restaurant

	const createRestaurant = `INSERT INTO restaurants (name, description, restaurant_id, owner_id, rating, created_at) values (?, ?, ?, ?, ?, ?)`

	switch r.Method {
	case "POST":
		err := json.NewDecoder(r.Body).Decode(&restaurant)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode("BadRequest")
			return
		}
		utils.OpenDb()
		utils.Db.Begin()
		_, err = utils.Db.Exec(
			createRestaurant,
			restaurant.Name,
			restaurant.Description,
			restaurant.RestaurantId,
			restaurant.OwnerId,
			restaurant.Rating, // TODO: ratingは、クライアント側が評価していく部分なので将来的にはcreateの際には設定できない様にする
			time.Now())
		if err != nil {
			log.Fatalln(err)
		} else {
			w.WriteHeader(201)
			json.NewEncoder(w).Encode("Created Restaurant")
		}
		defer utils.Db.Close()

		return
	}
}
