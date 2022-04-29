package user

import (
	"fmt"
	"go_api/pkg"
	"log"
	"time"
)

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	// create_at は、DB設定でparseTimeが必要だった
	const (
		createSession = `INSERT INTO sessions (uuid, email, user_id, created_at) values(?, ?, ?, ?)`
		selectSession = `SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = ? and email = ?`
	)

	pkg.OpenDb()
	result, err := pkg.Db.Exec(createSession, pkg.CreateUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)

	err = pkg.Db.QueryRow(selectSession, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt)
	if err != nil {
		log.Fatalln(err)
	}
	defer pkg.Db.Close()

	return
}
