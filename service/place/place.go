package place

import (
	"context"
	"net/url"

	"googlemaps.github.io/maps"
)

type ServiceGetPlacebyIDReq struct {
	PlaceID  string `json:"place_id"`
	Language string `json:"language"`
}

type ServiceGetPlacebyIDRes struct {
	PlaceID          string             `json:"place_id"`
	PlaceName        string             `json:"place_name"`
	FormattedAddress string             `json:"formatted_address"`
	Rating           float32            `json:"rating:"`
	UserRatingCount  int                `json:"user_rating_count"`
	PlaceType        []string           `json:"place_type"`
	Reviews          []maps.PlaceReview `json:"reviews"`
	GoogleMapsLink   string             `json:"google_maps_link"`
	Photos           []maps.Photo       `json:"photos"`
}

type Service struct {
	mapsClient *maps.Client
}

func NewService(mapsClient *maps.Client) *Service {
	return &Service{
		mapsClient: mapsClient,
	}
}

func (s *Service) GetByID(c context.Context, req *ServiceGetPlacebyIDReq) (*ServiceGetPlacebyIDRes, error) {
	searchQuery := &maps.PlaceDetailsRequest{
		PlaceID: req.PlaceID,
	}
	gmapResponse, err := s.mapsClient.PlaceDetails(c, searchQuery)
	if err != nil {
		return nil, err
	}

	urlEncodedName := url.QueryEscape(gmapResponse.Name)

	res := ServiceGetPlacebyIDRes{
		PlaceID:          gmapResponse.PlaceID,
		PlaceName:        gmapResponse.Name,
		FormattedAddress: gmapResponse.FormattedAddress,
		Rating:           gmapResponse.Rating,
		UserRatingCount:  gmapResponse.UserRatingsTotal,
		PlaceType:        gmapResponse.Types,
		Reviews:          gmapResponse.Reviews,
		GoogleMapsLink:   "https://www.google.com/maps/search/?api=1&query=" + urlEncodedName + "&query_place_id=" + gmapResponse.PlaceID,
		Photos:           gmapResponse.Photos,
	}
	return &res, err
}
