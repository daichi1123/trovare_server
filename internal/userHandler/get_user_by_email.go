package userHandler

import (
	"go_api/utils"
	"log"
)

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	const getUserByEmail = `SELECT id, uuid, name, email, password, created_at FROM users WHERE email = ?;`

	utils.OpenDb()
	err = utils.Db.QueryRow(getUserByEmail, email).Scan(&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt)
	if err != nil {
		log.Fatalln(err)
	}
	defer utils.Db.Close()

	return
}
