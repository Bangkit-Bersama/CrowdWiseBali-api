package main

import (
	"context"
	"net/http"

	firebase "firebase.google.com/go/v4"
	v1 "github.com/Bangkit-Bersama/CrowdWiseBali-api/internal/api/v1"
	"github.com/Bangkit-Bersama/CrowdWiseBali-api/internal/config"
	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/auth"
	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/place"
	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/prediction"
	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/recommendation"
	"github.com/Bangkit-Bersama/CrowdWiseBali-api/service/user"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"googlemaps.github.io/maps"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

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

	e.Validator = &CustomValidator{validator: validator.New()}

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

	authClient, err := firebaseClient.Auth(context.Background())
	if err != nil {
		e.Logger.Fatal(err)
	}

	firestoreClient, err := firebaseClient.Firestore(context.Background())
	if err != nil {
		e.Logger.Fatal(err)
	}

	mapsClient, err := maps.NewClient(maps.WithAPIKey(config.GMPAPIKey))
	if err != nil {
		e.Logger.Fatal(err)
	}

	authService := auth.NewService(authClient)
	userService := user.NewService(firestoreClient)
	placeService := place.NewService(mapsClient)
	recommendationService, err := recommendation.NewService(mapsClient)
	if err != nil {
		e.Logger.Fatal(err)
	}
	predictionService, err := prediction.NewService(firestoreClient)
	if err != nil {
		e.Logger.Fatal(err)
	}

	v1.NewGroup(
		e,
		authService,
		userService,
		placeService,
		recommendationService,
		predictionService,
	)

	e.Logger.Fatal(e.Start(":8080"))
}
