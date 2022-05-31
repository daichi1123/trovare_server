package restaurant

import (
	"encoding/json"
	"errors"
	"go_api/db"
	"go_api/query"
	"go_api/utils"
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
		pathParam := strings.TrimPrefix(r.URL.Path, "/v1/restaurants/")
		specifyID, err := strconv.Atoi(pathParam)
		if err != nil {
			log.Fatalln(err)
		}
		specify := SpecifyRestaurant{specifyID}

		db.OpenDb()
		db.Db.Begin()
		err = db.Db.QueryRow(query.ShowRst, specify.ID).Scan(
			&restaurant.ID,
			&restaurant.Name,
			&restaurant.Description,
			&restaurant.Rating,
			&restaurant.ZipCode,
			&restaurant.Address,
			&restaurant.ImageURL,
			&restaurant.CreatedAt)
		if err != nil {
			utils.ErrorJSON(w, errors.New("BadRequest"))
		}
		defer db.Db.Close()
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
