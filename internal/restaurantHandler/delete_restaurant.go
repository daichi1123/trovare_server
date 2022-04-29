package restaurantHandler

import (
	"errors"
	"go_api/utils"
	"net/http"
	"path"
	"strconv"
)

type specifyRestaurant struct {
	ID int `json:"id"`
}

func DeleteRestaurant(w http.ResponseWriter, r *http.Request) {
	const deleteRestaurant = `DELETE FROM restaurants WHERE id = ?;`

	switch r.Method {
	case http.MethodDelete:
		getID, err := strconv.Atoi(path.Base(r.URL.Path))
		if err != nil {
			w.WriteHeader(400)
			return
		}
		r := specifyRestaurant{getID}

		utils.OpenDb()
		utils.Db.Begin()
		stmt, err := utils.Db.Prepare(deleteRestaurant)
		if err != nil {
			utils.ErrorJSON(w, errors.New("BadRequest"))
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
		defer utils.Db.Close()
	}
}
