package query

const (
	GetSess          = `SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = ?`
	CreateSess       = `INSERT INTO sessions (uuid, email, user_id, created_at) values(?, ?, ?, ?)`
	SelectSess       = `SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = ? and email = ?`
	DeleteSessByUuid = `DELETE FROM sessions WHERE uuid = ?`
)
