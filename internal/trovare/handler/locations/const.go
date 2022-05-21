package locations

const (
	getFiveKmRadiusOfCurrentLocation = `SELECT id, name, description, zip_code, address,rating, lng, lat,
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
