package genre

import (
	"encoding/json"
	"fmt"
	"go_api/db"
	"go_api/query"
	"net/http"
	"time"
)

func CreateGenre(w http.ResponseWriter, r *http.Request) {
	var req CreateGenreRequest

	switch r.Method {
	case "POST":
		fmt.Println(&r.Body)
		err := json.NewDecoder(r.Body).Decode(&req)
		fmt.Println(&req)
		fmt.Println(req)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode("BadRequest")
			return
		}
		db.OpenDb()
		db.Db.Begin()
		_, err = db.Db.Exec(
			query.CreateG,
			req.Name,
			time.Now())
		defer db.Db.Close()
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode("BadRequest")
		} else {
			w.WriteHeader(201)
			json.NewEncoder(w).Encode("Created Genre")
		}
		w.Header().Set("Content-Type", "application/json")

		return
	}
}
