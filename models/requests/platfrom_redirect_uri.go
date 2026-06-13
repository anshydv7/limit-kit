package requests

import (
	"auth-service/constants"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type PlatformRedirectUriRequest struct {
	Platform    string `form:"platform" binding:"required"`
	RedirectUri string `form:"redirect_uri" binding:"required"`
	State       string `form:"state"`
}

func (req *PlatformRedirectUriRequest) Intitiate(c *gin.Context) (context.Context, error) {
	_ctx, _ := c.Get("context")
	ctx, _ := _ctx.(context.Context)
	if err := c.BindQuery(&req); err != nil {
		fmt.Println("error while binding request", err)
		return nil, err
	}

	return ctx, nil
}

func (req *PlatformRedirectUriRequest) Validate(ctx context.Context) error {
	req.Platform = strings.ToLower(strings.TrimSpace(req.Platform))
	if _, ok := constants.OauthPlatform[req.Platform]; !ok {
		return errors.New("invalid platform")
	}

	return nil
}
