package user

import (
	"fmt"
	"go_api/db"
	"go_api/query"
	"net/http"
	"time"
)

type Session struct {
	ID        int       `json:"id"`
	UUID      string    `json:"uuid"`
	Email     string    `json:"email"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

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
	db.OpenDb()

	err = db.Db.QueryRow(query.GetSess, sess.UUID).Scan(
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
	defer db.Db.Close()

	return
}
