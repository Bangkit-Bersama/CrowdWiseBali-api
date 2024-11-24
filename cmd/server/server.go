package main

import (
	"log"
	"net/http"

	v1Handler "github.com/Bangkit-Bersama/CrowdWiseBali-api/internal/api/v1"
	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/place"
	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/recommendation"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

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
