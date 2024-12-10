package v1

import (
	"net/http"
	"strings"

	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/auth"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	Service *auth.Service
}

func NewAuthHandler(g *echo.Group, service *auth.Service) *AuthHandler {
	handler := &AuthHandler{
		Service: service,
	}

	return handler
}

func (h *AuthHandler) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeaderRaw := c.Request().Header["Authorization"]
		if len(authHeaderRaw) == 0 {
			return echo.NewHTTPError(http.StatusUnauthorized, "token needed")
		}

		authHeader := authHeaderRaw[0]

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return echo.NewHTTPError(http.StatusUnauthorized, "token must be type Bearer")
		}

		token := authHeader[len("Bearer "):]

		_, err := h.Service.VerifyToken(c.Request().Context(), token)
		if err != nil {
			return err
		}

		return next(c)
	}
}
