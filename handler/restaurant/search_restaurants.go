package restaurant

import (
	"encoding/json"
	"fmt"
	"go_api/db"
	"go_api/models"
	"go_api/query"
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

		db.OpenDb()
		db.Db.Begin()

		if restaurant != "" {
			rows, err := db.Db.Query(query.SearchResult, restaurant)
			result := models.SearchRestaurantResponse{}
			var data []models.SearchRestaurantResponse

			for rows.Next() {
				err := rows.Scan(
					&result.Name,
					&result.Description,
					&result.Rating,
					&result.ZipCode,
					&result.Address,
					&result.Lat,
					&result.Lng,
					&result.ImageURL,
					&result.GenreName,
				)

				if err != nil {
					log.Println(err)
					w.WriteHeader(401)
				} else {
					data = append(data, result)
				}
			}
			defer rows.Close()

			rsts, err := json.Marshal(data)
			if err != nil {
				fmt.Println(err.Error())
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(rsts)
			return
		}

		if genre != "" {
			rows, err := db.Db.Query(query.SearchResultByGenre, genre)
			result := models.SearchRestaurantResponse{}
			var data []models.SearchRestaurantResponse

			for rows.Next() {
				err := rows.Scan(
					&result.Name,
					&result.Description,
					&result.Rating,
					&result.ZipCode,
					&result.Address,
					&result.Lat,
					&result.Lng,
					&result.ImageURL,
					&result.GenreName,
				)

				if err != nil {
					log.Println(err)
					w.WriteHeader(401)
				} else {
					data = append(data, result)
				}

				if err != nil {
					log.Fatal(err.Error())
				}
			}
			defer rows.Close()

			rstsByGenre, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(400)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(rstsByGenre)
			return
		}

		if rating != "" {
			rating, err := strconv.Atoi(rating)
			if err != nil {
				log.Println(err)
			}

			rows, err := db.Db.Query(query.SearchResultByRating, rating)
			result := models.SearchRestaurantResponse{}
			var data []models.SearchRestaurantResponse

			for rows.Next() {
				err := rows.Scan(
					&result.Name,
					&result.Description,
					&result.Rating,
					&result.ZipCode,
					&result.Address,
					&result.Lat,
					&result.Lng,
					&result.ImageURL,
					&result.GenreName,
				)

				if err != nil {
					log.Println(err)
					w.WriteHeader(401)
				} else {
					data = append(data, result)
				}

				if err != nil {
					log.Fatal(err.Error())
				}
			}
			defer rows.Close()

			rstsByRating, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(400)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(rstsByRating)
			return
		}
	}
	defer db.Db.Close()

	return
}
