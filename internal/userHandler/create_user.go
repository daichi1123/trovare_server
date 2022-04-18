package userHandler

import (
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
	const insert = `INSERT INTO users (uuid, name, email, password, created_at) values (?, ?, ?, ?, ?)`

	utils.OpenDb()

	utils.Db.Begin()
	_, err = utils.Db.Exec(
		insert,
		utils.CreateUUID(),
		u.Name,
		u.Email,
		utils.Encrypt(u.Password),
		time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	defer utils.Db.Close()

	return
}
