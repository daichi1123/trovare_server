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
	case http.MethodPost:
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
			restaurant.ZipCode,
			restaurant.Address,
			time.Now())
		defer pkg.Db.Close()
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
