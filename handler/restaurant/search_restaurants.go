package restaurant

import (
	"encoding/json"
	"fmt"
	"go_api/pkg"
	"log"
	"net/http"
	"strconv"
)

func SearchRestaurants(w http.ResponseWriter, r *http.Request) {
	// TODO: I have to add codes and query contents
	switch r.Method {
	case "GET":
		restaurant := r.URL.Query().Get("restaurants")
		genre := r.URL.Query().Get("genres")
		rating := r.URL.Query().Get("rating")

		pkg.OpenDb()
		pkg.Db.Begin()

		if restaurant != "" {
			rows, err := pkg.Db.Query(searchResult, restaurant)

			for rows.Next() {
				var name string
				err := rows.Scan(&name)

				if err != nil {
					log.Fatal(err.Error())
				}
				fmt.Println(name)
			}
			defer rows.Close()

			resp, err := json.Marshal(rows)
			if err != nil {
				w.WriteHeader(400)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(resp)
			return
		}

		if genre != "" {
			rows, err := pkg.Db.Query(searchResultByGenre, genre)

			for rows.Next() {
				var name string
				err := rows.Scan(&name)

				if err != nil {
					log.Fatal(err.Error())
				}
				fmt.Println(name)
			}
			defer rows.Close()

			restaurantByGenre, err := json.Marshal(rows)
			if err != nil {
				w.WriteHeader(400)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(restaurantByGenre)
			return
		}

		if rating != "" {
			rating, err := strconv.Atoi(rating)
			if err != nil {
				log.Println(err)
			}

			rows, err := pkg.Db.Query(searchResultByRating, rating)

			for rows.Next() {
				var name string
				err := rows.Scan(&name)

				if err != nil {
					log.Fatal(err.Error())
				}
				fmt.Println(name)
			}
			defer rows.Close()

			restaurantByRating, err := json.Marshal(rows)
			if err != nil {
				w.WriteHeader(400)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(restaurantByRating)
			return

		}
	}
	defer pkg.Db.Close()

	return
}
