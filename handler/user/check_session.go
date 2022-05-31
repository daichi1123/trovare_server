package user

import (
	"fmt"
	"go_api/db"
	"go_api/query"
)

func (sess *Session) CheckSession() (valid bool, err error) {
	fmt.Println(sess.UUID)
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
