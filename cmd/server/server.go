package main

import (
	"context"
	"net/http"

	// firebase "firebase.google.com/go"
	// "google.golang.org/api/option"
	"cloud.google.com/go/firestore"
	v1 "github.com/Bangkit-Bersama/CrowdWiseBali-api/internal/api/v1"
	"github.com/Bangkit-Bersama/CrowdWiseBali-api/internal/config"
	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/place"
	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/recommendation"
	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"googlemaps.github.io/maps"
)

func main() {
	config.LoadEnv()

	e := echo.New()
	if !config.Production {
		e.Logger.SetLevel(log.DEBUG)
		e.Debug = true
	}
	e.HideBanner = true

	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level}")
	}

	if !config.Production {
		e.Logger.Warn("Server is running in debug mode.")
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "API is running.")
	})

	mapsClient, err := maps.NewClient(maps.WithAPIKey(config.GMPAPIKey))
	if err != nil {
		e.Logger.Fatal(err)
	}

	fireCtx := context.Background()
	firestoreClient, err := firestore.NewClient(fireCtx, "crowdwise-bali")

	placeService := place.NewService(mapsClient)
	recommendationService := recommendation.NewService(mapsClient)
	userService := user.NewService(firestoreClient)

	v1.NewGroup(
		e,
		placeService,
		recommendationService,
		userService,
	)

	e.Logger.Fatal(e.Start(":8080"))
}
