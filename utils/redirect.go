package utils

import "net/http"

func RedirectIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/index", 302)
}
