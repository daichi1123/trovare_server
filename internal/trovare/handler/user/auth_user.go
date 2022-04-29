package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_api/pkg"
	"html/template"
	"log"
	"net/http"
	"path"
)

const tempDir = "templates"

func Login(w http.ResponseWriter, r *http.Request) {
	// ここでJWT or Sessionを使用してフロントに値を返すようにしたい
	t, err := template.ParseFiles(path.Join(tempDir + "/login.html"))
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, "")
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		fmt.Println("r", r)
		var creds Credentials
		// r.Bodyを　creds構造体の形でデコードする
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			pkg.ErrorJSON(w, errors.New("unauthorized"))
			return
		}

		//　emailからユーザ情報を受け取ることは成功
		// test用Email: test1@example.com
		user, err := GetUserByEmail(creds.Email)
		if err != nil {
			pkg.ErrorJSON(w, errors.New("unauthorized"))
			return
		}

		if user.Password == pkg.Encrypt(creds.Password) {
			// test用Password: test
			session, err := user.CreateSession()
			fmt.Print("session", session)
			if err != nil {
				log.Fatalln(err)
			}

			pkg.WriteJSON(w, http.StatusOK, session, "res")
		}

		//if user.Password != utils.Encrypt(creds.Password) {
		//	// test用Password: test
		//	session, err := user.CreateSession()
		//	fmt.Print("session", session)
		//	if err != nil {
		//		log.Fatalln(err)
		//	}
		//
		//	utils.WriteJSON(w, http.StatusOK, session, "res")
		//}
	}
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
	http.Redirect(w, r, "v1/login", 302)
}

func Authentication(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Fatalln(err)
		http.Redirect(w, r, "login", 302)
	}
	if user.Password == pkg.Encrypt(r.PostFormValue("password")) {
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
