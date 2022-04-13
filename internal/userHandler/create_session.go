package userHandler

import (
	"fmt"
	"go_api/utils"
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

	utils.OpenDb()
	result, err := utils.Db.Exec(createSession, utils.CreateUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)

	err = utils.Db.QueryRow(selectSession, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt)
	if err != nil {
		log.Fatalln(err)
	}
	defer utils.Db.Close()

	return
}
