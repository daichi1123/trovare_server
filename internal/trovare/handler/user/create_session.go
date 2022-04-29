package user

import (
	"fmt"
	"go_api/pkg"
	"log"
	"time"
)

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}

	pkg.OpenDb()
	result, err := pkg.Db.Exec(createSess, pkg.CreateUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)

	err = pkg.Db.QueryRow(selectSess, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt)
	if err != nil {
		log.Fatalln(err)
	}
	defer pkg.Db.Close()

	return
}
