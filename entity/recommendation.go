package entity

import (
	"googlemaps.github.io/maps"
)

type ServiceGetRecommendationReq struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	PlaceType string  `json:"place_type"`
}

type ServiceGetRecommendationRes struct {
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
