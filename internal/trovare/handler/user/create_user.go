package user

import (
	"go_api/pkg"
	"log"
	"time"
)

var createUser = User{
	Name:     "test1",
	Email:    "test1@example.com",
	Password: "test",
}

func (u User) CreateUser() (err error) {
	pkg.OpenDb()

	pkg.Db.Begin()
	_, err = pkg.Db.Exec(
		createU,
		pkg.CreateUUID(),
		u.Name,
		u.Email,
		pkg.Encrypt(u.Password),
		time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	defer pkg.Db.Close()

	return
}
