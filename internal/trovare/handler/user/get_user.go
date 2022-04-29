package user

import (
	"go_api/pkg"
	"log"
)

func GetUser(id int) (user User, err error) {
	user = User{}
	pkg.OpenDb()
	err = pkg.Db.QueryRow(getU, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password)
	if err != nil {
		log.Fatalln(err)
	}
	defer pkg.Db.Close()

	return
}
