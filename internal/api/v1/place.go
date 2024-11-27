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

	ng := g.Group("/places")

	ng.GET("/:id", handler.GetByID)

	return handler
}

func (h *PlaceHandler) GetByID(c echo.Context) error {
	return Response(c, http.StatusOK, map[string]string{
		"test": "hello",
	}, nil)
}
