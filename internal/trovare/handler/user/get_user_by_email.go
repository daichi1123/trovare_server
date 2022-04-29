package user

import (
	"fmt"
	"go_api/pkg"
	"log"
)

func GetUserByEmail(email string) (user User, err error) {
	user = User{}

	pkg.OpenDb()
	err = pkg.Db.QueryRow(getUByEmail, email).Scan(&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt)
	if err != nil {
		//TODO: このエラーの際frontのbodyにエラー分を返却したい
		log.Fatalln(err)
	}
	defer pkg.Db.Close()

	return
}
func GetUserByEmailForSession(email string) (user User, err error) {
	user = User{}
	const getUserByEmail = `SELECT id, uuid, name, email, password, created_at FROM users WHERE email = ?;`
	fmt.Print(email)

	pkg.OpenDb()
	err = pkg.Db.QueryRow(getUserByEmail, email).Scan(&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt)
	if err != nil {
		log.Fatalln(err)
	}
	defer pkg.Db.Close()

	return
}
