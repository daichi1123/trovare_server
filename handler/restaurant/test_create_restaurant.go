package restaurant

import (
	"go_api/db"
	"go_api/query"
	"log"
	"time"
)

func (r Restaurant) CreateTestRestaurant() (err error) {
	var cr = Restaurant{
		Name:        "cafe restaurant",
		Description: "You can eat delicious foods",
		Rating:      1,
		ZipCode:     1410031,
		Address:     "Tokyo shinagawaku nishigotanda 7-8",
		ImageURL:    "../../../../../img/curry.jpeg",
	}
	db.OpenDb()
	db.Db.Begin()
	_, err = db.Db.Exec(
		query.CreateRst,
		cr.Name,
		cr.Description,
		cr.Rating,
		cr.ZipCode,
		cr.Address,
		cr.ImageURL,
		time.Now())
	if err != nil {
		log.Println(err)
	}
	defer db.Db.Close()

	return
}
