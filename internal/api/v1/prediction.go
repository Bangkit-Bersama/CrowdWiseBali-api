package v1

import (
	"net/http"

	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/prediction"
	"github.com/labstack/echo/v4"
)

type PredictionRes struct {
	Occupancy float32 `json:"occupancy"`
}

type PredictionHandler struct {
	Service *prediction.Service
}

func NewPredictionHandler(g *echo.Group, service *prediction.Service, authHandler *AuthHandler) *PredictionHandler {
	handler := &PredictionHandler{
		Service: service,
	}

	r := g.Group("/prediction")

	r.Use(authHandler.AuthMiddleware)

	r.POST("", handler.PostPredict)

	return handler
}

func (h *PredictionHandler) PostPredict(c echo.Context) error {
	var req struct {
		PlaceId string `json:"placeId" validate:"required"`
		Date    string `json:"date" validate:"required"`
		Hour    int    `json:"hour" validate:"required"`
	}

	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	inference, err := h.Service.Predict(c.Request().Context(), &prediction.Req{
		PlaceId: req.PlaceId,
		Date:    req.Date,
		Hour:    req.Hour,
	})
	if err != nil {
		return err
	}

	return respond(c, http.StatusOK, &PredictionRes{
		Occupancy: inference,
	}, nil)
}
