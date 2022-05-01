package genre

import (
	"encoding/json"
	"fmt"
	"go_api/pkg"
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
		pkg.OpenDb()
		pkg.Db.Begin()
		_, err = pkg.Db.Exec(
			create,
			req.Name,
			time.Now())
		defer pkg.Db.Close()
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
