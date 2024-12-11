package v1

import (
	"net/http"

	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Service *user.Service
}

func NewUserHandler(g *echo.Group, service *user.Service) *UserHandler {
	handler := &UserHandler{
		Service: service,
	}

	r := g.Group("/users")

	r.GET("/:id", handler.GetUser)

	return handler
}
func (h *UserHandler) GetUser(c echo.Context) error {

	return respond(c, http.StatusOK, "Server menyala abangkuh!!!", nil)
}
