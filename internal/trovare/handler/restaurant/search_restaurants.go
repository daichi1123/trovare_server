package restaurant

import (
	"encoding/json"
	"fmt"
	"go_api/pkg"
	"log"
	"net/http"
)

func SearchRestaurants(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.URL.Query() {

		}

		restaurant := r.URL.Query().Get("restaurants")
		log.Println(restaurant)

		pkg.OpenDb()
		pkg.Db.Begin()

		rows, err := pkg.Db.Query(getSearchResult, restaurant)
		defer rows.Close()
		defer pkg.Db.Close()

		for rows.Next() {
			var name string
			err := rows.Scan(&name)

			if err != nil {
				panic(err.Error())
			}
			fmt.Println(name)
		}

		//var rstsInfoRes []CurrentLocationRes
		//for rows.Next() {
		//	currentLocationRes := CurrentLocationRes{}
		//
		//	//ここの値をcurrentLocationRes.Location.Latみたいにすればいけるのではないか
		//	err = rows.Scan(
		//		&currentLocationRes.ID,
		//		&currentLocationRes.Name,
		//		&currentLocationRes.Description,
		//		&currentLocationRes.ZipCode,
		//		&currentLocationRes.Address,
		//		&currentLocationRes.Rating,
		//		&currentLocationRes.Coords.Lat,
		//		&currentLocationRes.Coords.Lng,
		//		&distance.Distance)
		//	rstsInfoRes = append(rstsInfoRes, currentLocationRes)
		//}
		//if err != nil {
		//	w.WriteHeader(401)
		//	log.Println(err)
		//	return
		//}
		//defer rows.Close()
		//defer pkg.Db.Close()
		//
		resp, err := json.Marshal(rows)
		if err != nil {
			w.WriteHeader(400)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return
	}
	return
}
