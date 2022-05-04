package restaurant

import (
	"encoding/json"
	"go_api/pkg"
	"net/http"
)

func GetAllRestaurants(w http.ResponseWriter, r *http.Request) {
	pkg.OpenDb()
	pkg.Db.Begin()
	rows, err := pkg.Db.Query(getAll)
	if err != nil {
		w.WriteHeader(500)
	}
	defer rows.Close()

	data := []Restaurant{}
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
	defer pkg.Db.Close()

	res, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// TODO: この処理をまとめておきたい
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}
