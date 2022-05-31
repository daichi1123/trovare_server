package user

import (
	"fmt"
	"go_api/db"
	"go_api/query"
	"go_api/utils"
	"log"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}

	db.OpenDb()
	result, err := db.Db.Exec(query.CreateSess, utils.CreateUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)

	err = db.Db.QueryRow(query.SelectSess, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Db.Close()

	return
}
