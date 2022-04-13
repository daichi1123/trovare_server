package pkg

import (
	"go_api/internal/auth"
	"html/template"
	"log"
	"net/http"
	"path"
)

const tempDir = "templates"

func Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		_, err := auth.ConfirmSession(w, r)
		if err != nil {
			t, err := template.ParseFiles(path.Join(tempDir + "/index.html"))
			if err != nil {
				log.Fatal(err)
			}
			t.Execute(w, "")
		} else {
			http.Redirect(w, r, "/after-login", 302)
		}
	default:
		log.Printf("ERROR: Invalid HTTP Method")
	}
}
