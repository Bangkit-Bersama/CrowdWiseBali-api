package v1

import (
	"net/http"

	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/recommendation"
	"github.com/labstack/echo/v4"
)

type RecommendationHandler struct {
	Service *recommendation.Service
}

func NewRecommendationHandler(g *echo.Group, service *recommendation.Service, authHandler *AuthHandler) *RecommendationHandler {
	handler := &RecommendationHandler{
		Service: service,
	}

	routeGroup := g.Group("/recommendation")

	routeGroup.Use(authHandler.AuthMiddleware)

	routeGroup.GET("", handler.GetByLocation)

	return handler
}

func (h *RecommendationHandler) GetByLocation(c echo.Context) error {
	var req struct {
		latitude  float64
		longitude float64
		placeType string
	}

	err := echo.QueryParamsBinder(c).
		MustFloat64("latitude", &req.latitude).
		MustFloat64("longitude", &req.longitude).
		MustString("place_type", &req.placeType).
		BindError()
	if err != nil {
		return err
	}

	recommendations, err := h.Service.GetByLocation(c.Request().Context(), &recommendation.ReqData{
		Latitude:  req.latitude,
		Longitude: req.longitude,
	})
	if err != nil {
		return err
	}

	return respond(c, http.StatusOK, recommendations, nil)
}
