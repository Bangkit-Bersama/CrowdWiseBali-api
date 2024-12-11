package v1

import (
	"net/http"

	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/auth"
	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/place"
	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/recommendation"
	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/user"
	"github.com/labstack/echo/v4"
)

type V1Group struct {
	e *echo.Echo

	route *echo.Group

	auth           *AuthHandler
	user           *UserHandler
	place          *PlaceHandler
	recommendation *RecommendationHandler
	prediction     *PredictionHandler
}

func NewGroup(
	e *echo.Echo,
	authService *auth.Service,
	userService *user.Service,
	placeService *place.Service,
	recommendationService *recommendation.Service,
) *V1Group {
	route := e.Group("/api/v1")

	route.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)

			if err != nil {
				if he, ok := err.(*echo.HTTPError); ok {
					return respond(c, he.Code, nil, err)
				}

				if be, ok := err.(*echo.BindingError); ok {
					return respond(c, http.StatusBadRequest, map[string]interface{}{
						be.Field: be.Message,
					}, nil)
				}

				c.Logger().Error(err)
				return respond(c, http.StatusInternalServerError, nil, nil)
			}

			return nil
		}
	})

	authHandler := NewAuthHandler(route, authService)

	v1Group := &V1Group{
		e:              e,
		route:          route,
		auth:           authHandler,
		user:           NewUserHandler(route, userService),
		place:          NewPlaceHandler(route, placeService, authHandler),
		recommendation: NewRecommendationHandler(route, recommendationService, authHandler),
		Echo:                  e,
		RouteGroup:            routeGroup,
		placeHandler:          NewPlaceHandler(routeGroup, placeService, authHandler),
		recommendationHandler: NewRecommendationHandler(routeGroup, recommendationService, authHandler),
		userHandler:           NewUserHandler(routeGroup, userService),
	}

	return v1Group
}

func respond(c echo.Context, code int, i interface{}, e error) error {
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
