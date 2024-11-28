package recommendation

import (
	"context"
	"net/url"

	"github.com/Bangkit-Bersama/CrowdWiseBali-api/entity"
	"github.com/Bangkit-Bersama/CrowdWiseBali-api/internal/config"
	"github.com/labstack/echo/v4"
	"googlemaps.github.io/maps"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetByLocation(c echo.Context, req entity.ServiceGetRecommendationReq) (entity.ServiceGetRecommendationRes, error) {
	res := entity.ServiceGetRecommendationRes{}

	apiKey := config.GMPAPIKey

	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return res, err
	}

	location := &maps.LatLng{
		Lat: req.Latitude,
		Lng: req.Longitude,
	}

	radius := 5000 //in meter

	searchQuery := &maps.NearbySearchRequest{
		Location: location,
		Radius:   uint(radius),
		Type:     maps.PlaceType(req.PlaceType),
	}

	gmapResponse, err := client.NearbySearch(context.Background(), searchQuery)
	if err != nil {
		return res, err
	}

	for _, result := range gmapResponse.Results {
		urlEncodedName := url.QueryEscape(result.Name)
		googleMapsLink := "https://www.google.com/maps/search/?api=1&query=" + urlEncodedName + "&query_place_id=" + result.PlaceID
		places := entity.NearbyPlaces{
			PlaceID:         result.PlaceID,
			PlaceName:       result.Name,
			UserRatingCount: result.UserRatingsTotal,
			Rating:          result.Rating,
			GoogleMapsLink:  googleMapsLink,
			Photos:          result.Photos,
		}
		res.SearchResult = append(res.SearchResult, places)
	}

	return res, nil
}
