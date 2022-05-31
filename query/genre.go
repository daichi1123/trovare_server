package query

const (
	GetAllG      = `SELECT * FROM genres`
	ShowG        = `SELECT id, name, description, rating, created_at FROM genres WHERE id = ?`
	CreateG      = `INSERT INTO genres (name, created_at) values (?, ?)`
	UpdateG      = `UPDATE genres SET name = ? WHERE id = ?`
	DeleteG      = `DELETE FROM genres WHERE id = ?`
	InsertGenres = `INSERT INTO genres(id, name, created_at) VALUES( ?, ?, ?)`
)
