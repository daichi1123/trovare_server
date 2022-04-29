package restaurantHandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_api/utils"
	"log"
	"net/http"
	"time"
)

// Dummy Restaurant
var DummyRestaurant = Restaurant{
	ID: 1,
	// そのままだとエラーになるmysqlはutf8で3バイトまでしか扱えないのでutf8mb4にする必要がある
	Name:         "ステーキハウス",
	Description:  "とても美味しいステーキ屋です。",
	RestaurantId: 3,
	OwnerId:      1,
	Rating:       4,
}

// w http.ResponseWriter, r *http.Request
func (restaurant Restaurant) RegisterRestaurant() (err error) {
	// ownerIDを作成する必要あり
	const (
		registerRestaurant = `INSERT INTO restaurants (name, description, restaurant_id, owner_id, rating, created_at) values (?, ?, ?, ?, ?, ?)`
	)

	utils.OpenDb()

	utils.Db.Begin()
	//rest, err := utils.Db.Exec(
	//	registerRestaurant,
	//	restaurant.Name,
	//	restaurant.Description,
	//	restaurant.RestaurantId,
	//	restaurant.OwnerId,
	//	restaurant.Rating,
	//	time.Now())
	fmt.Println(DummyRestaurant.Name)
	rest, err := utils.Db.Exec(
		registerRestaurant,
		DummyRestaurant.Name,
		DummyRestaurant.Description,
		DummyRestaurant.RestaurantId,
		DummyRestaurant.OwnerId,
		DummyRestaurant.Rating,
		time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("restaurantのテスト", rest)

	return
}

func RegisterRestaurantForWeb(w http.ResponseWriter, r *http.Request) {
	// frontからjsonの値を受け取り保存することに成功 frontに何か値を返す必要性があるのか
	var restaurant Restaurant

	const (
		registerRestaurant = `INSERT INTO restaurants (name, description, restaurant_id, owner_id, rating, created_at) values (?, ?, ?, ?, ?, ?)`
	)

	switch r.Method {
	case http.MethodPost:
		err := json.NewDecoder(r.Body).Decode(&restaurant)
		if err != nil {
			utils.ErrorJSON(w, errors.New("BadRequest"))
			return
		}

		utils.OpenDb()
		utils.Db.Begin()
		_, err = utils.Db.Exec(
			registerRestaurant,
			restaurant.Name,
			restaurant.Description,
			restaurant.RestaurantId,
			restaurant.OwnerId,
			restaurant.Rating, // ここは設定できてはいけないので0で設定しておくのがありかも
			time.Now())
		if err != nil {
			log.Fatalln(err)
		} else {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode("Created")
		}
		defer utils.Db.Close()

		return
	}
}
