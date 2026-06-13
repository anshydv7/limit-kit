package authfactory

import "context"

type GoogleAuth struct {
	clientId     string
	clientSecret string
}

func (req *GoogleAuth) OauthUrl(ctx context.Context, redirectUri, state string) {

}
