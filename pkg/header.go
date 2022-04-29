package pkg

import "net/http"

func HeaderInfo(w http.ResponseWriter, r http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	w.Header().Add("Content-Type", "application/json")
}
