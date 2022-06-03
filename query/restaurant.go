package query

const (
	GetAllRsts     = `SELECT * FROM restaurants`
	GetLists       = `SELECT id, name, description, zip_code, address, image_url, rating FROM restaurants`
	GetRstLocation = `SELECT id, name, zip_code, address, image_url, rating FROM restaurants`
	CreateRst      = `INSERT INTO restaurants (name, description, rating, zip_code, address, lat, lng, image_url, created_at) values (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	SearchResult   = `SELECT r.name, r.description, r.rating, r.zip_code, r.address, r.lat, r.lng, r.image_url, g.name
						FROM restaurants AS r JOIN genres AS g ON r.genre_id=g.id WHERE r.name LIKE CONCAT('%', ?, '%')`
	SearchResultByGenre = `SELECT r.name, r.description, r.rating, r.zip_code, r.address, r.lat, r.lng, r.image_url, g.name 
							FROM restaurants AS r JOIN genres AS g ON r.genre_id=g.id WHERE r.genre_id=?`
	SearchResultByRating = `SELECT r.name, r.description, r.rating, r.zip_code, r.address, r.lat, r.lng, r.image_url, g.name
							FROM restaurants AS r JOIN genres AS g ON r.genre_id=g.id WHERE r.rating=?`
	InsertRsts = `INSERT INTO restaurants (name, description, rating, zip_code, address, lat, lng, image_url, genre_id, created_at) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	ShowRst    = `SELECT id, name, description, rating, zip_code, address, image_url, created_at FROM restaurants WHERE id = ?`
	UpdateRst  = `UPDATE restaurants SET name = ? WHERE id = ?`
	DeleteRst  = `DELETE FROM restaurants WHERE id = ?`
)

const (
	GetFiveKmRadiusOfCurrentLocation = `SELECT id, name, description, zip_code, address,rating, lng, lat,
										(
											6378 * acos(
												cos( radians(?)) 
												* cos( radians( lat ) )
												* cos( radians( lng ) - radians(?) )
												+ sin( radians(?) )
												* sin( radians( lat ))
											)
										) AS distance
										FROM
											restaurants
										HAVING
											distance <= 5
										ORDER BY
											distance
										LIMIT 0 , 20
										;`
)
