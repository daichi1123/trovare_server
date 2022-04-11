package userHandler

import (
	"go_api/utils"
	"log"
)

func GetUser(id int) (user User, err error) {
	user = User{}
	const selectUser = `SELECT id, uuid, name, email, password FROM users WHERE id = ?`
	utils.OpenDb()
	err = utils.Db.QueryRow(selectUser, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password)
	if err != nil {
		log.Fatalln(err)
	}
	return
}
