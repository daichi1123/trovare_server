package auth

import (
	"fmt"
	"go_api/pkg"
	"html/template"
	"log"
	"net/http"
	"path"
)

func ConfirmSession(w http.ResponseWriter, r *http.Request) (sess Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = Session{UUID: cookie.Value}
		if ok, _ := sess.checkSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return
}

// 共通化したい
func (sess *Session) checkSession() (valid bool, err error) {
	fmt.Println(sess.UUID)
	pkg.OpenDb()
	const getSession = `SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = ?`

	err = pkg.Db.QueryRow(getSession, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.UserID,
		&sess.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if sess.ID != 0 {
		valid = true
	}
	defer pkg.Db.Close()

	return
}

func AfterLogin(w http.ResponseWriter, r *http.Request) {
	_, err := ConfirmSession(w, r)
	if err != nil {
		http.Redirect(w, r, "/index", 302)
	} else {
		t, err := template.ParseFiles(path.Join(tempDir + "/after_login.html"))
		if err != nil {
			log.Fatalln(err)
		}
		t.Execute(w, "")
	}
}
