package main

import (
	"net/http"

	v1Handler "github.com/Bangkit-Bersama/CrowdWiseBali-api/internal/api/v1"
	"github.com/Bangkit-Bersama/CrowdWiseBali-api/internal/config"
	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/place"
	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/recommendation"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func errorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)
	c.JSON(code, map[string]string{
		"status":  "error",
		"message": err.Error(),
	})
}

func main() {
	config.LoadEnv()

	e := echo.New()
	e.Debug = !config.Production
	e.HideBanner = true
	e.HTTPErrorHandler = errorHandler

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "API is running.")
	})

	placeService := place.NewService()
	recommendationService := recommendation.NewService()

	v1Group := e.Group("/api/v1")
	v1Handler.NewPlaceHandler(v1Group, placeService)
	v1Handler.NewRecommendationHandler(v1Group, recommendationService)

	log.Fatal(e.Start(":8080"))
}
