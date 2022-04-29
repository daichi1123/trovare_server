package restaurantHandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_api/utils"
	"net/http"
)

type specifyRestaurant struct {
	ID int `json:"id"`
}

func DeleteRestaurant(w http.ResponseWriter, r *http.Request) {
	var specifyID specifyRestaurant

	const deleteRestaurant = `DELETE FROM restaurants WHERE id = ?;`

	switch r.Method {
	case http.MethodDelete:
		fmt.Println(r.Body)
		err := json.NewDecoder(r.Body).Decode(&specifyID)
		if err != nil {
			utils.ErrorJSON(w, errors.New("BadRequest"))
			return
		}

		utils.OpenDb()
		utils.Db.Begin()
		stmt, err := utils.Db.Prepare(deleteRestaurant)
		if err != nil {
			utils.ErrorJSON(w, errors.New("BadRequest"))
			return
		}

		_, err = stmt.Exec(specifyID.ID)
		if err != nil {
			utils.ErrorJSON(w, errors.New("BadRequest"))
			return
		} else {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		defer utils.Db.Close()
	}
}
