package restaurant

import (
	"encoding/json"
	"errors"
	"go_api/pkg"
	"net/http"
	"path"
	"strconv"
)

type updateInfo struct {
	Name string `json:"name"`
}

func UpdateRestaurantInfo(w http.ResponseWriter, r *http.Request) {
	var updateRestaurantName updateInfo

	switch r.Method {
	case http.MethodPatch:
		getID, err := strconv.Atoi(path.Base(r.URL.Path))
		if err != nil {
			w.WriteHeader(400)
			return
		}

		err = json.NewDecoder(r.Body).Decode(&updateRestaurantName)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		r := SpecifyRestaurant{getID}

		pkg.OpenDb()
		pkg.Db.Begin()
		result, err := pkg.Db.Exec(update, updateRestaurantName.Name, r.ID)
		if err != nil {
			pkg.ErrorJSON(w, errors.New("BadRequest"))
		}
		defer pkg.Db.Close()
		json.Marshal(result)

		return
	}
	return
}
