package restaurant

import (
	"encoding/json"
	"go_api/pkg"
	"net/http"
	"time"
)

// name of method need `_'table name' or not`??

func CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	var restaurant Restaurant

	switch r.Method {
	case "POST":
		err := json.NewDecoder(r.Body).Decode(&restaurant)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode("BadRequest")
			return
		}
		pkg.OpenDb()
		pkg.Db.Begin()
		_, err = pkg.Db.Exec(
			create,
			restaurant.Name,
			restaurant.Description,
			restaurant.RestaurantId,
			restaurant.OwnerId,
			restaurant.Rating, // TODO: ratingは、クライアント側が評価していく部分なので将来的にはcreateの際には設定できない様にする
			time.Now())
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode("BadRequest")
		}
		w.WriteHeader(201)
		json.NewEncoder(w).Encode("Created Restaurant")
		defer pkg.Db.Close()

		return
	}
}
