package responses

import "auth-service/models/domain"

type AuthenticationResponse struct {
	User  *domain.User `json:"user"`
	Token string       `json:"token"`
}
