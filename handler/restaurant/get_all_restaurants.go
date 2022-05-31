package restaurant

import (
	"encoding/json"
	"go_api/db"
	"go_api/query"
	"net/http"
)

func GetAllRestaurants(w http.ResponseWriter, r *http.Request) {
	db.OpenDb()
	db.Db.Begin()
	rows, err := db.Db.Query(query.GetAllRsts)
	if err != nil {
		w.WriteHeader(500)
	}
	defer rows.Close()

	var data []Restaurant
	for rows.Next() {
		restaurant := Restaurant{}

		err = rows.Scan(
			&restaurant.ID,
			restaurant.Name,
			restaurant.Description,
			restaurant.ZipCode,
			restaurant.Address,
			&restaurant.Rating,
			&restaurant.CreatedAt,
		)
		data = append(data, restaurant)
	}
	defer rows.Close()

	if err != nil {
		w.WriteHeader(400)
	}
	defer db.Db.Close()

	res, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}
