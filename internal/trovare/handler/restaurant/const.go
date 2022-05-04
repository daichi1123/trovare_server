package restaurant

const (
	getAll = `SELECT * FROM restaurants`
	create = `INSERT INTO restaurants (name, description, zip_code, address, created_at) values (?, ?, ?, ?, ?)`
	show   = `SELECT id, name, description, rating, created_at FROM restaurants WHERE id = ?`
	update = `UPDATE restaurants SET name = ? WHERE id = ?`
	delete = `DELETE FROM restaurants WHERE id = ?`
)
