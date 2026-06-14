package authfactory

import (
	"auth-service/models/domain"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type GoogleAuth struct {
	clientId     string
	clientSecret string
}

func (ga *GoogleAuth) GetOauthUrl(ctx context.Context, redirectUri, state string) string {
	u, err := url.Parse("https://accounts.google.com/o/oauth2/v2/auth")
	if err != nil {
		return ""
	}
	q := u.Query()
	q.Set("client_id", ga.clientId)
	q.Set("redirect_uri", redirectUri)
	q.Set("response_type", "code")
	q.Set("scope", "https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/userinfo.email")
	q.Set("state", state)
	u.RawQuery = q.Encode()
	return u.String()
}

func (ga *GoogleAuth) Authentication(ctx context.Context, redirectUri, code string) (*domain.GoogleUser, error) {
	accessToken, err := ga.getAccessToken(ctx, redirectUri, code)
	if err != nil {
		return nil, err
	}
	return ga.getUserInfo(ctx, accessToken)
}

func (ga *GoogleAuth) getAccessToken(ctx context.Context, redirectUri, code string) (string, error) {
	tokenUrl := "https://oauth2.googleapis.com/token"
	formValues := url.Values{}
	formValues.Set("client_id", ga.clientId)
	formValues.Set("client_secret", ga.clientSecret)
	formValues.Set("code", code)
	formValues.Set("grant_type", "authorization_code")
	formValues.Set("redirect_uri", redirectUri)

	req, err := http.NewRequestWithContext(ctx, "POST", tokenUrl, strings.NewReader(formValues.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create token request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send token request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("token request failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var tokenResponse struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return "", fmt.Errorf("failed to decode token response: %w", err)
	}

	if tokenResponse.AccessToken == "" {
		return "", fmt.Errorf("received empty access token")
	}

	return tokenResponse.AccessToken, nil
}

func (ga *GoogleAuth) getUserInfo(ctx context.Context, accessToken string) (*domain.GoogleUser, error) {
	userInfoUrl := "https://www.googleapis.com/oauth2/v2/userinfo"
	reqInfo, err := http.NewRequestWithContext(ctx, "GET", userInfoUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create user info request: %w", err)
	}
	reqInfo.Header.Set("Authorization", "Bearer "+accessToken)

	respInfo, err := http.DefaultClient.Do(reqInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to send user info request: %w", err)
	}
	defer respInfo.Body.Close()

	if respInfo.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(respInfo.Body)
		return nil, fmt.Errorf("user info request failed with status %d: %s", respInfo.StatusCode, string(bodyBytes))
	}

	var userInfo domain.GoogleUser
	if err := json.NewDecoder(respInfo.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode user info response: %w", err)
	}

	return &userInfo, nil
}
