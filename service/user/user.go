package user

import (
	"context"

	"cloud.google.com/go/firestore"
)

type Service struct {
	firestoreClient *firestore.Client
}

func NewService(firestoreClient *firestore.Client) *Service {
	return &Service{
		firestoreClient: firestoreClient,
	}
}

func (s *Service) GetUser(c context.Context) error {
	// users := s.firestoreClient.Collection("users")
	// // user := users.Doc()
	// // docSnap, err := users.Parent.Get()

	return nil
}
