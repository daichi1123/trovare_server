package service

type (
	Distance struct {
		Distance float64
	}

	CurrentLocationRequest struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}

	CurrentLocationResponse struct {
		ID          int     `json:"id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		ZipCode     int     `json:"zip_code"`
		Address     string  `json:"address"`
		Rating      int     `json:"Rating"`
		Lat         float64 `json:"lat"`
		Lng         float64 `json:"lng"`
	}
)
