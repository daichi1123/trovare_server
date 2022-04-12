package pkg

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

const tempDir = "templates"

func Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		t, err := template.ParseFiles(path.Join(tempDir + "/index.html"))
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, "")
	default:
		log.Printf("ERROR: Invalid HTTP Method")
	}
}
