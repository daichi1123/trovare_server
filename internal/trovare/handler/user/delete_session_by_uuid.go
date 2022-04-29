package user

import (
	"go_api/pkg"
	"log"
)

func (sess *Session) DeleteSessionByUUID() (err error) {
	pkg.OpenDb()
	_, err = pkg.Db.Exec(deleteSessByUuid, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	defer pkg.Db.Close()

	return
}
