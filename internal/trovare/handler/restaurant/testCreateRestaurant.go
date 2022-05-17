package restaurant

import (
	"go_api/pkg"
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
	pkg.OpenDb()
	pkg.Db.Begin()
	_, err = pkg.Db.Exec(
		create,
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
	defer pkg.Db.Close()

	return
}
