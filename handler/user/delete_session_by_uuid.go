package user

import (
	"go_api/db"
	"go_api/query"
	"log"
)

func (sess *Session) DeleteSessionByUUID() (err error) {
	db.OpenDb()
	_, err = db.Db.Exec(query.DeleteSessByUuid, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Db.Close()

	return
}
