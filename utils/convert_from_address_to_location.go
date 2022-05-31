package utils

import (
	"context"
	"flag"
	"go_api/configs"
	"googlemaps.github.io/maps"
	"log"
)

type Location struct {
	Lat float64 `json:"prefecture"`
	Lng float64 `json:"locality"`
}

var (
	language = flag.String("language", "ja", "The language in which to return results.")
	region   = flag.String("region", "JP", "The region code, specified as a ccTLD two-character value.")
)

func ConvertFromAddressToLocation(address string) Location {
	flag.Parse()

	var client *maps.Client
	var err error

	client, err = maps.NewClient(maps.WithAPIKey(configs.Config.GoogleMapsApiKey))
	outputFatalLog(err)

	req := &maps.GeocodingRequest{
		Address:  address,
		Language: *language,
		Region:   *region,
	}
	resp, err := client.Geocode(context.Background(), req)
	outputFatalLog(err)

	location := Location{
		Lat: resp[0].Geometry.Location.Lat,
		Lng: resp[0].Geometry.Location.Lng,
	}

	return Location{Lat: location.Lat, Lng: location.Lng}
}

func outputFatalLog(err error) {
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
}
