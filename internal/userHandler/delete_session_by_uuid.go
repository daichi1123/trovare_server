package userHandler

import (
	"go_api/utils"
	"log"
)

func (sess *Session) DeleteSessionByUUID() (err error) {
	const deleteSessionByUuid = `DELETE FROM sessions WHERE uuid = ?`
	utils.OpenDb()
	_, err = utils.Db.Exec(deleteSessionByUuid, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	defer utils.Db.Close()

	return
}
