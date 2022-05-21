package locations

import (
	"encoding/json"
	"log"
	"net/http"
)

func PostCurrentLocation(w http.ResponseWriter, r *http.Request) {
	var currentLocation CurrentLocationReq

	switch r.Method {
	case http.MethodPost:
		err := json.NewDecoder(r.Body).Decode(&currentLocation)
		if err != nil {
			w.WriteHeader(401)
			log.Println(err)
			return
		}

		log.Println(currentLocation) // NOTE:現在地の取得に成功
		return
	}
}
