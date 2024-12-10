package v1

import (
	"net/http"

	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/place"
	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/recommendation"
	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/user"
	"github.com/labstack/echo/v4"
)

type V1Group struct {
	Echo *echo.Echo

	RouteGroup *echo.Group

	placeHandler          *PlaceHandler
	recommendationHandler *RecommendationHandler
	userHandler           *UserHandler
}

func NewGroup(
	e *echo.Echo,
	userService *user.Service,
	placeService *place.Service,
	recommendationService *recommendation.Service,
) *V1Group {
	routeGroup := e.Group("/api/v1")

	routeGroup.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)

			if err != nil {
				if he, ok := err.(*echo.HTTPError); ok {
					return Respond(c, he.Code, nil, err)
				}

				if be, ok := err.(*echo.BindingError); ok {
					return Respond(c, http.StatusBadRequest, map[string]interface{}{
						be.Field: be.Message,
					}, nil)
				}

				c.Logger().Error(err)
				return Respond(c, http.StatusInternalServerError, nil, nil)
			}

			return nil
		}
	})

	v1Group := &V1Group{
		Echo:                  e,
		RouteGroup:            routeGroup,
		placeHandler:          NewPlaceHandler(routeGroup, placeService),
		recommendationHandler: NewRecommendationHandler(routeGroup, recommendationService),
		userHandler:           NewUserHandler(routeGroup, userService),
	}

	return v1Group
}

func Respond(c echo.Context, code int, i interface{}, e error) error {
	response := map[string]interface{}{
		"status": "success",
	}

	if i != nil {
		response["data"] = i
	}

	if e != nil {
		response["message"] = e.Error()
	}

	if code >= 400 && code <= 499 {
		response["status"] = "fail"
	} else if code >= 500 && code <= 599 {
		response["status"] = "error"
	}

	return c.JSON(code, response)
}
