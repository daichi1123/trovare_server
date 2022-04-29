package restaurant

import (
	"encoding/json"
	"errors"
	"go_api/pkg"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type SpecifyRestaurant struct {
	ID int `json:"id"`
}

func GetRestaurant(w http.ResponseWriter, r *http.Request) {
	var restaurant Restaurant

	switch r.Method {
	case http.MethodGet:
		pathParam := strings.TrimPrefix(r.URL.Path, "/v1/restaurant/")
		specifyID, err := strconv.Atoi(pathParam)
		if err != nil {
			log.Fatalln(err)
		}
		specify := SpecifyRestaurant{specifyID}

		pkg.OpenDb()
		pkg.Db.Begin()
		err = pkg.Db.QueryRow(show, specify.ID).Scan(
			&restaurant.ID,
			&restaurant.Name,
			&restaurant.Description,
			&restaurant.Rating,
			&restaurant.CreatedAt)
		if err != nil {
			pkg.ErrorJSON(w, errors.New("BadRequest"))
		}
		defer pkg.Db.Close()
		res, err := json.Marshal(restaurant)
		if err != nil {
			w.WriteHeader(400)
		}

		w.Header().Set("Location", r.Host+r.URL.Path+strconv.Itoa(specify.ID))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(res)

		return
	}
}
