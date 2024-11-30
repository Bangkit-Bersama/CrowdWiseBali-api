package recommendation

import (
	"context"
	"net/url"

	"googlemaps.github.io/maps"
)

type ReqData struct {
	Latitude  float64
	Longitude float64
}

type ResData struct {
	SearchResult []NearbyPlaces `json:"search_result"`
}

type NearbyPlaces struct {
	PlaceID         string       `json:"place_id"`
	PlaceType       []string     `json:"place_type"`
	PlaceName       string       `json:"place_name"`
	UserRatingCount int          `json:"user_rating_count"`
	Rating          float32      `json:"rating"`
	GoogleMapsLink  string       `json:"google_maps_link"`
	Photos          []maps.Photo `json:"photos"`
}

type Service struct {
	mapsClient *maps.Client
}

func NewService(mapsClient *maps.Client) *Service {
	return &Service{
		mapsClient: mapsClient,
	}
}

func (s *Service) GetByLocation(c context.Context, req *ReqData) (*ResData, error) {
	location := &maps.LatLng{
		Lat: req.Latitude,
		Lng: req.Longitude,
	}

	radius := 10000 //in meter

	searchQuery := &maps.NearbySearchRequest{
		Location: location,
		Radius:   uint(radius),
	}

	gmapResponse, err := s.mapsClient.NearbySearch(c, searchQuery)
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
			PlaceType:       result.Types,
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
