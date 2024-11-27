package v1

import (
	"net/http"
	"strconv"

	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/recommendation"
	"github.com/labstack/echo/v4"
)

type RecommendationHandler struct {
	Service *recommendation.Service
}

func NewRecommendationHandler(g *echo.Group, service *recommendation.Service) *RecommendationHandler {
	handler := &RecommendationHandler{
		Service: service,
	}

	ng := g.Group("/recommendation")

	ng.GET("", handler.GetByLocation)

	return handler
}

func (h *RecommendationHandler) GetByLocation(c echo.Context) error {
	latitudeParam := c.QueryParam("latitude")
	longitudeParam := c.QueryParam("longitude")
	placeType := c.QueryParam("placeType")

	invalids := make([]string, 0, 4)

	latitude, err := strconv.ParseFloat(latitudeParam, 64)
	if err != nil {
		invalids = append(invalids, "Invalid latitude")
		// Response(c, http.StatusBadRequest, "Invalid latitude", nil)
	}

	longitude, err := strconv.ParseFloat(longitudeParam, 64)
	if err != nil {
		invalids = append(invalids, "Invalid longitude")
		// Response(c, http.StatusBadRequest, "Invalid longitude", nil)
	}

	if placeType == "" {
		invalids = append(invalids, "placeType is required")
	}

	if len(invalids) > 0 {
		return Response(c, http.StatusBadRequest, invalids, nil)
	}

	recommendations, err := h.Service.GetByLocation(c, latitude, longitude, placeType)
	if err != nil {
		return err
	}

	return Response(c, http.StatusOK, recommendations, nil)
}
