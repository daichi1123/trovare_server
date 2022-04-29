package user

import (
	"fmt"
	"go_api/pkg"
)

func (sess *Session) CheckSession() (valid bool, err error) {
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
