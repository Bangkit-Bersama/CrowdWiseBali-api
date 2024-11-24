package v1

import (
	"errors"

	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/place"
	"github.com/labstack/echo/v4"
)

type PlaceHandler struct {
	Service *place.Service
}

func NewPlaceHandler(g *echo.Group, service *place.Service) *PlaceHandler {
	handler := &PlaceHandler{
		Service: service,
	}

	ng := g.Group("/places")

	ng.GET("/:id", handler.GetByID)

	return handler
}

func (h *PlaceHandler) GetByID(c echo.Context) error {
	return errors.New("not implemented")
}
