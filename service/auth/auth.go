package auth

import (
	"context"

	"firebase.google.com/go/v4/auth"
)

type Service struct {
	authClient *auth.Client
}

func NewService(authClient *auth.Client) *Service {
	return &Service{
		authClient: authClient,
	}
}

func (s *Service) VerifyToken(context context.Context, token string) (parsedToken *auth.Token, err error) {
	tk, err := s.authClient.VerifyIDTokenAndCheckRevoked(context, token)
	if err != nil {
		return nil, err
	}

	return tk, nil
}
