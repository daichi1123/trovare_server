package restaurantHandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_api/utils"
	"net/http"
	"path"
	"strconv"
)

type updateInfo struct {
	Name string `json:"name"`
}

func UpdateRestaurantInfo(w http.ResponseWriter, r *http.Request) {
	var restaurant Restaurant
	var updateRestaurantName updateInfo

	//const fileを作成してその中で変数の管理をしたい
	const updateRestaurant = `UPDATE restaurants SET name = ? WHERE id = ?`

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

		utils.OpenDb()
		utils.Db.Begin()
		_, err = utils.Db.Exec(updateRestaurant, updateRestaurantName.Name, r.ID)
		if err != nil {
			utils.ErrorJSON(w, errors.New("BadRequest"))
		}
		defer utils.Db.Close()
		json.Marshal(restaurant)
		fmt.Print(restaurant)

		return
	}
	return
}
