package user

import (
	"go_api/db"
	"go_api/query"
	"log"
)

func GetUserByEmail(email string) (user User, err error) {
	user = User{}

	db.OpenDb()
	err = db.Db.QueryRow(query.GetUByEmail, email).Scan(&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt)
	if err != nil {
		//TODO: このエラーの際frontのbodyにエラー分を返却したい
		log.Fatalln(err)
	}
	defer db.Db.Close()

	return
}
func GetUserByEmailForSession(email string) (user User, err error) {
	user = User{}
	db.OpenDb()
	err = db.Db.QueryRow(query.GetUByEmail, email).Scan(&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Db.Close()

	return
}
