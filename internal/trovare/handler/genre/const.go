package genre

const (
	getAll = `SELECT * FROM genres`
	show   = `SELECT id, name, description, rating, created_at FROM genres WHERE id = ?`
	create = `INSERT INTO genres (name, created_at) values (?, ?)`
	update = `UPDATE genres SET name = ? WHERE id = ?`
	delete = `DELETE FROM genres WHERE id = ?`
)
