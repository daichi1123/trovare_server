package query

// created_at は、DB設定でparseTimeが必要だった
const (
	GetU        = `SELECT id, uuid, name, email, password FROM users WHERE id = ?`
	CreateU     = `INSERT INTO users (uuid, name, email, password, created_at) values (?, ?, ?, ?, ?)`
	GetUByEmail = `SELECT id, uuid, name, email, password, created_at FROM users WHERE email = ?;`
)
