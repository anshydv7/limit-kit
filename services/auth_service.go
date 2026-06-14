package services

import (
	"auth-service/models/domain"
	"auth-service/models/requests"
	"auth-service/models/responses"
	"auth-service/repository"
	authfactory "auth-service/services/auth_factory"
	"auth-service/utils"
	"context"
	"database/sql"
)

func GetOauthUrl(ctx context.Context, request *requests.PlatformRedirectUriRequest) (string, error) {
	authFactory, err := authfactory.NewAuthFactory(request.Platform)
	if err != nil {
		return "", err
	}
	url := authFactory.GetOauthUrl(ctx, request.RedirectUri, request.State)
	return url, nil
}

func Authenticate(ctx context.Context, request *requests.AuthenticationRequest) (*responses.AuthenticationResponse, error) {
	authFactory, err := authfactory.NewAuthFactory(request.Platform)
	if err != nil {
		return nil, err
	}

	userData, err := authFactory.Authentication(ctx, request.RedirectUri, request.Code)
	if err != nil {
		return nil, err
	}

	user, err := repository.GetUserByEmail(ctx, userData.Email)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}

		newUser := &domain.User{
			Email:          userData.Email,
			Name:           userData.Name,
			ProfilePicture: userData.Picture,
		}

		user, err = repository.CreateUser(ctx, newUser)
		if err != nil {
			return nil, err
		}
	}

	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return &responses.AuthenticationResponse{
		User:  user,
		Token: token,
	}, nil
}
