package restaurantHandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_api/utils"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type SpecifyRestaurant struct {
	ID int `json:"id"`
}

func GetRestaurant(w http.ResponseWriter, r *http.Request) {
	var restaurant Restaurant

	//const fileを作成してその中で変数の管理をしたい
	const getRestaurant = `SELECT id, name, description, rating, created_at FROM restaurants WHERE id = ?`

	switch r.Method {
	case http.MethodGet:
		pathParam := strings.TrimPrefix(r.URL.Path, "/v1/restaurant/")
		specifyID, err := strconv.Atoi(pathParam)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(reflect.TypeOf(specifyID))

		specify := SpecifyRestaurant{specifyID}
		fmt.Println(specify)

		utils.OpenDb()
		utils.Db.Begin()
		err = utils.Db.QueryRow(getRestaurant, specify.ID).Scan(
			&restaurant.ID,
			&restaurant.Name,
			&restaurant.Description,
			&restaurant.Rating,
			&restaurant.CreatedAt)
		if err != nil {
			utils.ErrorJSON(w, errors.New("BadRequest"))
		}
		defer utils.Db.Close()
		json.Marshal(restaurant)

		return
	}
	return
}
