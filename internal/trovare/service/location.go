package service

type CurrentLocationRequest struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// NOTE: 住所を登録する際は、緯度経度で登録する方が後々使い勝手が良さそう
