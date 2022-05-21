package locations

import (
	"encoding/json"
	"log"
	"net/http"
)

// TODO: lat lngをフィルターで半径5キロ分のデータを取得できる様にする
// TODO: そこで取得できた情報をフロントエンドに送れる様にする

func PostCurrentLocation(w http.ResponseWriter, r *http.Request) {
	// WIP
	var currentLocation CurrentLocationReq

	switch r.Method {
	case http.MethodPost:
		err := json.NewDecoder(r.Body).Decode(&currentLocation)
		log.Println(getFiveKmRadiusOfCurrentLocation)
		if err != nil {
			w.WriteHeader(401)
			log.Println(err)
			return
		}

		log.Println(currentLocation) // NOTE:現在地の取得に成功
		return
	}
}
