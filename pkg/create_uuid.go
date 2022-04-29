package pkg

import "github.com/google/uuid"

func CreateUUID() (uuidobj uuid.UUID) {
	// uuidの作成
	uuidobj, _ = uuid.NewUUID()
	return
}
