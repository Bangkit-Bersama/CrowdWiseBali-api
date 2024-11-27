package recommendation

import (
	"github.com/labstack/echo/v4"
	"googlemaps.github.io/maps"
)

type Service interface {
	GetByLocation(c echo.Context, latitude float64, longitude float64, placeType string) (maps.PlacesSearchResponse, error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}
