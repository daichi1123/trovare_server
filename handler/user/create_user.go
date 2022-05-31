package user

import (
	"go_api/db"
	"go_api/query"
	"go_api/utils"
	"log"
	"time"
)

var createUser = User{
	Name:     "test1",
	Email:    "test1@example.com",
	Password: "test",
}

func (u User) CreateUser() (err error) {
	db.OpenDb()

	db.Db.Begin()
	_, err = db.Db.Exec(
		query.CreateU,
		utils.CreateUUID(),
		u.Name,
		u.Email,
		utils.Encrypt(u.Password),
		time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Db.Close()

	return
}
