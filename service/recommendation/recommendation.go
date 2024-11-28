package recommendation

import (
	"context"
	"net/url"

	"github.com/Bangkit-Bersama/CrowdWiseBali-api/internal/config"
	"googlemaps.github.io/maps"
)

type ReqData struct {
	Latitude  float64
	Longitude float64
	PlaceType string
}

type ResData struct {
	SearchResult []NearbyPlaces `json:"search_result"`
}

type NearbyPlaces struct {
	PlaceID         string       `json:"place_id"`
	PlaceName       string       `json:"place_name"`
	UserRatingCount int          `json:"user_rating_count"`
	Rating          float32      `json:"rating"`
	GoogleMapsLink  string       `json:"google_maps_link"`
	Photos          []maps.Photo `json:"photos"`
}

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetByLocation(req *ReqData) (*ResData, error) {
	apiKey := config.GMPAPIKey

	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
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
		return nil, err
	}

	res := ResData{
		SearchResult: make([]NearbyPlaces, 0, len(gmapResponse.Results)),
	}

	for _, result := range gmapResponse.Results {
		urlEncodedName := url.QueryEscape(result.Name)
		googleMapsLink := "https://www.google.com/maps/search/?api=1&query=" + urlEncodedName + "&query_place_id=" + result.PlaceID
		places := NearbyPlaces{
			PlaceID:         result.PlaceID,
			PlaceName:       result.Name,
			UserRatingCount: result.UserRatingsTotal,
			Rating:          result.Rating,
			GoogleMapsLink:  googleMapsLink,
			Photos:          result.Photos,
		}
		res.SearchResult = append(res.SearchResult, places)
	}

	return &res, nil
}
