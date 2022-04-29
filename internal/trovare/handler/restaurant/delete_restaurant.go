package restaurant

import (
	"errors"
	"go_api/pkg"
	"net/http"
	"path"
	"strconv"
)

type specifyRestaurant struct {
	ID int `json:"id"`
}

func DeleteRestaurant(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		getID, err := strconv.Atoi(path.Base(r.URL.Path))
		if err != nil {
			w.WriteHeader(400)
			return
		}
		r := specifyRestaurant{getID}

		pkg.OpenDb()
		pkg.Db.Begin()
		stmt, err := pkg.Db.Prepare(delete)
		if err != nil {
			pkg.ErrorJSON(w, errors.New("BadRequest"))
			return
		}

		_, err = stmt.Exec(r.ID)
		if err != nil {
			w.WriteHeader(400)
			return
		} else {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		defer pkg.Db.Close()
	}
}
