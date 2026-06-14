package authfactory

import (
	"auth-service/base/config"
	"auth-service/models/domain"
	"context"
	"errors"
)

type AuthFactory interface {
	GetOauthUrl(ctx context.Context, redirectUri, state string) string
	Authentication(ctx context.Context, redirectUri, code string) (*domain.GoogleUser, error)
}

func NewAuthFactory(platform string) (AuthFactory, error) {
	switch platform {
	case "google":
		return &GoogleAuth{
			clientId:     config.Config.Google.ClientID,
			clientSecret: config.Config.Google.ClientSecret,
		}, nil
	default:
		return nil, errors.New("invalid platform")
	}
}
