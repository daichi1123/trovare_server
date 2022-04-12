package userHandler

import (
	"go_api/utils"
	"html/template"
	"log"
	"net/http"
	"path"
)

const tempDir = "templates"

func Login(w http.ResponseWriter, r *http.Request) {
	//_, err := session(w, r)
	t, err := template.ParseFiles(path.Join(tempDir + "/login.html"))
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, "")
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Fatal(err)
	}

	if err != http.ErrNoCookie {
		session := Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
	}
	http.Redirect(w, r, "/login", 302)
}

func Authentication(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Fatalln(err)
		http.Redirect(w, r, "login", 302)
	}
	if user.Password == utils.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Fatalln(err)
		}

		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/index", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}
