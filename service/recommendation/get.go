package recommendation

import (
	"context"
	"log"

	"github.com/Bangkit-Bersama/CrowdWiseBali-api/config"

	"github.com/labstack/echo/v4"
	"googlemaps.github.io/maps"
)

func (s *service) GetByLocation(c echo.Context, latitude float64, longitude float64, placeType string) (maps.PlacesSearchResponse, error) {
	res := maps.PlacesSearchResponse{}

	cfg, err := config.LoadConfig()

	apiKey := cfg.GMPKey

	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return res, err
	}

	location := &maps.LatLng{
		Lat: latitude,
		Lng: longitude,
	}

	radius := 5000 //in meter

	req := &maps.NearbySearchRequest{
		Location: location,
		Radius:   uint(radius),
		Type:     maps.PlaceType(placeType),
	}

	res, err = client.NearbySearch(context.Background(), req)
	if err != nil {
		return res, err
	}

	log.Println("Nearby Places:")
	for _, result := range res.Results {
		log.Printf("PlaceID: %s, Name: %s, Address: %s, Rating: %.1f",
			result.PlaceID, result.Name, result.Vicinity, result.Rating)
	}

	return res, nil
}
