package recommendation

import (
	"context"
	"encoding/json"
	"net/url"
	"os"

	"googlemaps.github.io/maps"
)

type Req struct {
	Latitude  float64
	Longitude float64
}

type Res struct {
	SearchResult []NearbyPlaces `json:"search_result"`
}

type NearbyPlaces struct {
	PlaceID          string       `json:"place_id"`
	PlaceType        []string     `json:"place_type"`
	PlaceName        string       `json:"place_name"`
	FormattedAddress string       `json:"formatted_address"`
	UserRatingCount  int          `json:"user_rating_count"`
	Rating           float32      `json:"rating"`
	GoogleMapsLink   string       `json:"google_maps_link"`
	Photos           []maps.Photo `json:"photos"`
}

type Service struct {
	mapsClient *maps.Client

	trainedPlaceIds []string
}

func NewService(mapsClient *maps.Client) (*Service, error) {
	placesJsonRaw, err := os.ReadFile("./places.json")
	if err != nil {
		return nil, err
	}

	var placesJson struct {
		PlaceIds map[string]string `json:"placeIds"`
	}

	if err = json.Unmarshal([]byte(placesJsonRaw), &placesJson); err != nil {
		return nil, err
	}

	var trainedPlaceIds []string
	for k := range placesJson.PlaceIds {
		trainedPlaceIds = append(trainedPlaceIds, k)
	}

	return &Service{
		mapsClient:      mapsClient,
		trainedPlaceIds: trainedPlaceIds,
	}, nil
}

func (s *Service) GetByLocation(c context.Context, req *Req) (*Res, error) {
	// For now, ignore location and only serve 9 places the ML model has trained to.

	count := len(s.trainedPlaceIds)

	resChan := make(chan *maps.PlaceDetailsResult, count)

	ctx, cancel := context.WithCancelCause(c)

	for i := 0; i < count; i++ {
		go func(idx int) {
			if ctx.Err() != nil {
				return
			}

			// fmt.Printf("Requesting place id %s...\n", s.trainedPlaceIds[i])

			searchQuery := &maps.PlaceDetailsRequest{
				PlaceID: s.trainedPlaceIds[i],
			}
			gmapResponse, err := s.mapsClient.PlaceDetails(ctx, searchQuery)
			if err != nil {
				cancel(err)
				return
			}

			resChan <- &gmapResponse
		}(i)
	}

	res := &Res{
		SearchResult: make([]NearbyPlaces, 0, count),
	}

	for i := 0; i < count; i++ {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				return nil, err
			}
		case result := <-resChan:
			urlEncodedName := url.QueryEscape(result.Name)
			googleMapsLink := "https://www.google.com/maps/search/?api=1&query=" + urlEncodedName + "&query_place_id=" + result.PlaceID
			res.SearchResult = append(res.SearchResult, NearbyPlaces{
				PlaceID:          result.PlaceID,
				PlaceType:        result.Types,
				PlaceName:        result.Name,
				FormattedAddress: result.FormattedAddress,
				UserRatingCount:  result.UserRatingsTotal,
				Rating:           result.Rating,
				GoogleMapsLink:   googleMapsLink,
				Photos:           result.Photos,
			})
		}
	}

	return res, nil
}

/*
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
			PlaceID:          result.PlaceID,
			PlaceType:        result.Types,
			PlaceName:        result.Name,
			FormattedAddress: result.FormattedAddress,
			UserRatingCount:  result.UserRatingsTotal,
			Rating:           result.Rating,
			GoogleMapsLink:   googleMapsLink,
			Photos:           result.Photos,
		}
		res.SearchResult = append(res.SearchResult, places)
	}

	return &res, nil
}
*/
