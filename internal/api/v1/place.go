package v1

import (
	"net/http"

	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/place"
	"github.com/labstack/echo/v4"
)

type PlaceHandler struct {
	Service *place.Service
}

func NewPlaceHandler(g *echo.Group, service *place.Service, authHandler *AuthHandler) *PlaceHandler {
	handler := &PlaceHandler{
		Service: service,
	}

	routeGroup := g.Group("/places")

	routeGroup.Use(authHandler.AuthMiddleware)

	routeGroup.GET("/:id", handler.GetByID)

	return handler
}

func (h *PlaceHandler) GetByID(c echo.Context) error {
	var req struct {
		placeID  string
		language string
	}

	err := echo.PathParamsBinder(c).
		MustString("id", &req.placeID).
		BindError()
	if err != nil {
		return err
	}

	err = echo.QueryParamsBinder(c).
		MustString("language", &req.language).
		BindError()
	if err != nil {
		return err
	}

	places, err := h.Service.GetByID(c.Request().Context(), &place.ServiceGetPlacebyIDReq{
		PlaceID:  req.placeID,
		Language: req.language,
	})
	if err != nil {
		return err
	}

	return respond(c, http.StatusOK, places, nil)
}
