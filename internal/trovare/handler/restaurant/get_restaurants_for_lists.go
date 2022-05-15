package restaurant

import (
	"encoding/json"
	"go_api/internal/trovare/service"
	"go_api/pkg"
	"log"
	"net/http"
)

func GetRestaurantsForList(w http.ResponseWriter, r *http.Request) {
	pkg.OpenDb()
	pkg.Db.Begin()
	rows, err := pkg.Db.Query(getLists)
	if err != nil {
		w.WriteHeader(500)
	}
	defer rows.Close()

	restaurant := service.GetRestaurantsListsResponse{}
	var data []service.GetRestaurantsListsResponse

	for rows.Next() {
		err = rows.Scan(
			&restaurant.ID,
			&restaurant.Name,
			&restaurant.Description,
			&restaurant.ZipCode,
			&restaurant.Address,
			&restaurant.ImageURL,
			&restaurant.Rating)
		if err != nil {
			log.Println(err)
			w.WriteHeader(401)
		} else {
			data = append(data, restaurant)
		}
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

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}
