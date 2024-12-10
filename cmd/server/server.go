package main

import (
	"context"
	"net/http"

	// firebase "firebase.google.com/go"
	// "google.golang.org/api/option"

	firebase "firebase.google.com/go/v4"
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

	firebaseClient, err := firebase.NewApp(context.Background(), &firebase.Config{
		ProjectID: "crowdwise-bali",
	})
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Logger.Debug(firebaseClient)

	firestoreClient, err := firebaseClient.Firestore(context.Background())
	if err != nil {
		e.Logger.Fatal(err)
	}

	mapsClient, err := maps.NewClient(maps.WithAPIKey(config.GMPAPIKey))
	if err != nil {
		e.Logger.Fatal(err)
	}

	userService := user.NewService(firestoreClient)
	placeService := place.NewService(mapsClient)
	recommendationService := recommendation.NewService(mapsClient)

	v1.NewGroup(
		e,
		userService,
		placeService,
		recommendationService,
	)

	e.Logger.Fatal(e.Start(":8080"))
}
