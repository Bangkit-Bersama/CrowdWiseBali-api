package v1

import (
	"net/http"

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

	routeGroup := g.Group("/places")

	routeGroup.GET("/:id", handler.GetByID)

	return handler
}

func (h *PlaceHandler) GetByID(c echo.Context) error {
	return Respond(c, http.StatusOK, map[string]string{
		"test": "hello",
	}, nil)
}
