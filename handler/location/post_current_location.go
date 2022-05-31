package location

import (
	"encoding/json"
	"go_api/db"
	"go_api/query"
	"log"
	"net/http"
)

func PostCurrentLocation(w http.ResponseWriter, r *http.Request) {
	var currentLocation CurrentLocationReq

	switch r.Method {
	case http.MethodPost:
		err := json.NewDecoder(r.Body).Decode(&currentLocation)

		db.OpenDb()
		db.Db.Begin()
		rows, err := db.Db.Query(query.GetFiveKmRadiusOfCurrentLocation,
			currentLocation.Lat,
			currentLocation.Lng,
			currentLocation.Lat)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
		}

		var rstsInfoRes []CurrentLocationRes
		log.Println(CurrentLocationRes{})
		for rows.Next() {
			currentLocationRes := CurrentLocationRes{}
			// TODO: 下記structを使用せずにquery文を実行したい
			distance := Distance{}

			//ここの値をcurrentLocationRes.Location.Latみたいにすればいけるのではないか
			err = rows.Scan(
				&currentLocationRes.ID,
				&currentLocationRes.Name,
				&currentLocationRes.Description,
				&currentLocationRes.ZipCode,
				&currentLocationRes.Address,
				&currentLocationRes.Rating,
				&currentLocationRes.Coords.Lat,
				&currentLocationRes.Coords.Lng,
				&distance.Distance)
			rstsInfoRes = append(rstsInfoRes, currentLocationRes)
		}
		if err != nil {
			w.WriteHeader(401)
			log.Println(err)
			return
		}
		defer rows.Close()
		defer db.Db.Close()

		resp, err := json.Marshal(rstsInfoRes)
		if err != nil {
			w.WriteHeader(400)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return
	}
}
