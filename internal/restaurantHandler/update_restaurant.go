package restaurantHandler

import (
	"encoding/json"
	"errors"
	"go_api/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
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
		pathParam := strings.TrimPrefix(r.URL.Path, "/v1/restaurant/update/")
		specifyID, err := strconv.Atoi(pathParam)
		if err != nil {
			log.Fatalln(err)
		}

		err = json.NewDecoder(r.Body).Decode(&updateRestaurantName)
		if err != nil {
			log.Fatalln(err)
		}

		specify := SpecifyRestaurant{specifyID}

		utils.OpenDb()
		utils.Db.Begin()
		// updateInfoNameに問題がある
		result, err := utils.Db.Exec(updateRestaurant, updateRestaurantName.Name, specify.ID)
		if err != nil {
			utils.ErrorJSON(w, errors.New("BadRequest"))
		}
		defer utils.Db.Close()
		json.Marshal(restaurant)
		json.Marshal(result)

		return
	}
	return
}
