package user

import (
	"go_api/db"
	"go_api/query"
	"log"
)

func GetUser(id int) (user User, err error) {
	user = User{}
	db.OpenDb()
	err = db.Db.QueryRow(query.GetU, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Db.Close()

	return
}
