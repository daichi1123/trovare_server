package user

const (
	// user
	getU        = `SELECT id, uuid, name, email, password FROM users WHERE id = ?`
	createU     = `INSERT INTO users (uuid, name, email, password, created_at) values (?, ?, ?, ?, ?)`
	getUByEmail = `SELECT id, uuid, name, email, password, created_at FROM users WHERE email = ?;`

	// session
	getSess = `SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = ?`
	// created_at は、DB設定でparseTimeが必要だった
	createSess       = `INSERT INTO sessions (uuid, email, user_id, created_at) values(?, ?, ?, ?)`
	selectSess       = `SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = ? and email = ?`
	deleteSessByUuid = `DELETE FROM sessions WHERE uuid = ?`
)
