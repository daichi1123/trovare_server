package user

import (
	"go_api/pkg"
	"log"
)

func (sess *Session) DeleteSessionByUUID() (err error) {
	const deleteSessionByUuid = `DELETE FROM sessions WHERE uuid = ?`
	pkg.OpenDb()
	_, err = pkg.Db.Exec(deleteSessionByUuid, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	defer pkg.Db.Close()

	return
}
