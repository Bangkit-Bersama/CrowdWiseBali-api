package prediction

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/Bangkit-Bersama/CrowdWiseBali-api/internal/config"
	"google.golang.org/api/idtoken"
)

type Req struct {
	PlaceId string `json:"placeId"`
	Date    string `json:"date"`
	Hour    int    `json:"hour"`
}

type inferenceRes struct {
	Occupancy float32 `json:"prediction"`
}

type Service struct {
	firestoreClient *firestore.Client
	inferenceClient *http.Client
}

func NewService(firestoreClient *firestore.Client) (*Service, error) {
	reqClient, err := idtoken.NewClient(context.Background(), config.InferenceServerUrl)
	if err != nil {
		return nil, err
	}

	return &Service{
		firestoreClient: firestoreClient,
		inferenceClient: reqClient,
	}, nil
}

func (s *Service) Predict(ctx context.Context, req *Req) (float32, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return 0, err
	}

	res, err := s.inferenceClient.Post(config.InferenceServerUrl, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return 0, errors.New("inference server error")
	}

	inference := &inferenceRes{}
	err = json.NewDecoder(res.Body).Decode(inference)
	if err != nil {
		return 0, err
	}

	return inference.Occupancy, nil
}
