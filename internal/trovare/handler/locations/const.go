package locations

const (
	getFiveKmRadiusOfCurrentLocation = `SELECT id, name, description, rating, lng, lat, (6378 * acos(cos(radians(35.668185)) * cos(radians(lat)) * cos(radians(lng) – radians(139.739487)) + sin(radians(35.668185)) * sin(radians(lat)))) AS distance FROM restaurants HAVING distance < 10 ORDER BY distance`
)
