package restaurant

import (
	"encoding/json"
	"go_api/db"
	"go_api/query"
	"go_api/utils"
	"net/http"
	"time"
)

// name of method need `_'table name' or not`??

func CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	var restaurant Restaurant

	switch r.Method {
	case http.MethodPost:
		err := json.NewDecoder(r.Body).Decode(&restaurant)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode("BadRequest")
			return
		}
		location := utils.ConvertFromAddressToLocation(restaurant.Address)

		db.OpenDb()
		db.Db.Begin()
		_, err = db.Db.Exec(
			query.CreateRst,
			restaurant.Name,
			restaurant.Description,
			restaurant.Rating,
			restaurant.ZipCode,
			restaurant.Address,
			location.Lat,
			location.Lng,
			restaurant.ImageURL,
			time.Now())
		defer db.Db.Close()
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode("BadRequest")
		} else {
			w.WriteHeader(201)
			json.NewEncoder(w).Encode("Created Restaurant")
		}
		return
	}
}
