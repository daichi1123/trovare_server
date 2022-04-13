package userHandler

import (
	"fmt"
	"go_api/utils"
)

func (sess *Session) CheckSession() (valid bool, err error) {
	fmt.Println(sess.UUID)
	utils.OpenDb()
	const getSession = `SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = ?`

	err = utils.Db.QueryRow(getSession, sess.UUID).Scan(
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
	defer utils.Db.Close()

	return
}
