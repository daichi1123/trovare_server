package restaurant

import (
	"encoding/json"
	"go_api/pkg"
	"log"
	"net/http"
)

func GetRstLocation(w http.ResponseWriter, r *http.Request) {
	pkg.OpenDb()
	pkg.Db.Begin()
	rows, err := pkg.Db.Query(getRstLocation)
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
	defer pkg.Db.Close()

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
