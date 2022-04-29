package restaurantHandler

import (
	"fmt"
	"go_api/utils"
	"log"
	"net/http"
)

//WIP:
func GetAllRestaurants(w http.ResponseWriter, r *http.Request) {
	const indexRestaurants = `SELECT * FROM restaurants`

	utils.OpenDb()
	utils.Db.Begin()
	rows, err := utils.Db.Query(indexRestaurants)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	fmt.Println(rows)
	var restaurant *Restaurant
	fmt.Println("ssss")
	err = rows.Scan(
		&restaurant.ID,
		&restaurant.Name,
		&restaurant.Description,
		&restaurant.RestaurantId,
		&restaurant.OwnerId,
		&restaurant.Rating,
		&restaurant.CreatedAt,
	)
	if err != nil {
		log.Fatal(err)
	}
	//for rows.Next() {
	//}
	defer utils.Db.Close()
	//json.marshal()
}
