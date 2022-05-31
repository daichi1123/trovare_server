package restaurant

import (
	"encoding/json"
	"go_api/db"
	"go_api/query"
	"log"
	"net/http"
)

func GetRstLocation(w http.ResponseWriter, r *http.Request) {
	db.OpenDb()
	db.Db.Begin()
	rows, err := db.Db.Query(query.GetRstLocation)
	if err != nil {
		log.Println(err)
		w.WriteHeader(401)
	}

	rstLocation := GetRestaurantLocationRequest{}
	var data []GetRestaurantLocationRequest

	for rows.Next() {
		err = rows.Scan(
			&rstLocation.ID,
			&rstLocation.Name,
			&rstLocation.ZipCode,
			&rstLocation.Address,
			&rstLocation.ImageURL)
		if err != nil {
			log.Println(err)
			w.WriteHeader(401)
			return
		} else {
			data = append(data, rstLocation)
		}
	}
	defer rows.Close()

	res, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}
	defer db.Db.Close()

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
