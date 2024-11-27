package v1

import (
	"net/http"
	"strconv"

	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/recommendation"
	"github.com/labstack/echo/v4"
)

type RecommendationHandler struct {
	Service recommendation.Service
}

func NewRecommendationHandler(g *echo.Group, service recommendation.Service) *RecommendationHandler {
	handler := &RecommendationHandler{
		Service: service,
	}

	ng := g.Group("/recommendation")

	ng.GET("/", handler.GetByLocation)

	return handler
}

func (h *RecommendationHandler) GetByLocation(c echo.Context) error {
	latitudeParam := c.QueryParam("latitude")
	longitudeParam := c.QueryParam("longitude")
	placeType := c.QueryParam("placeType")

	latitude, err := strconv.ParseFloat(latitudeParam, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid latitude"})
	}

	longitude, err := strconv.ParseFloat(longitudeParam, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid longitude"})
	}

	if placeType == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "placeType is required"})
	}

	recommendations, err := h.Service.GetByLocation(c, latitude, longitude, placeType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, recommendations)
}
