package auth

import "net/http"

func HandleRestaurantsRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
	case "POST":
	case "PATCH":
	case "DELETE":
	default:
		w.WriteHeader(405)
	}
}
