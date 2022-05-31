package restaurant

import (
	"encoding/json"
	"go_api/db"
	"go_api/models"
	"go_api/query"
	"log"
	"net/http"
)

func GetRestaurantsForList(w http.ResponseWriter, r *http.Request) {
	db.OpenDb()
	db.Db.Begin()
	rows, err := db.Db.Query(query.GetLists)
	if err != nil {
		w.WriteHeader(500)
	}
	defer rows.Close()

	restaurant := models.GetRestaurantsListsResponse{}
	var data []models.GetRestaurantsListsResponse

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
