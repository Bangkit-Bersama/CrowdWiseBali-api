package recommendation

import (
	"context"
	"log"

	"github.com/Bangkit-Bersama/CrowdWiseBali-api/internal/config"
	"github.com/labstack/echo/v4"
	"googlemaps.github.io/maps"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetByLocation(c echo.Context, latitude float64, longitude float64, placeType string) (maps.PlacesSearchResponse, error) {
	res := maps.PlacesSearchResponse{}

	apiKey := config.GMPAPIKey

	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return res, err
	}

	location := &maps.LatLng{
		Lat: latitude,
		Lng: longitude,
	}

	radius := 5000 // in meter

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
