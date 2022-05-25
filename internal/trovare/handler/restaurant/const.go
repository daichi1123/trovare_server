package restaurant

const (
	getAll          = `SELECT * FROM restaurants`
	getLists        = `SELECT id, name, description, zip_code, address, image_url, rating FROM restaurants`
	getRstLocation  = `SELECT id, name, zip_code, address, image_url, rating FROM restaurants`
	getSearchResult = `SELECT name FROM restaurants WHERE name LIKE CONCAT('%', ?, '%')`
	create          = `INSERT INTO restaurants (name, description, rating, zip_code, address, lat, lng, image_url, created_at) values (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	show            = `SELECT id, name, description, rating, zip_code, address, image_url, created_at FROM restaurants WHERE id = ?`
	update          = `UPDATE restaurants SET name = ? WHERE id = ?`
	delete          = `DELETE FROM restaurants WHERE id = ?`
)
