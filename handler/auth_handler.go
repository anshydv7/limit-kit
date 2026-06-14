package handler

import (
	"auth-service/models/requests"
	"auth-service/models/responses"
	"auth-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOauthUrl(c *gin.Context) {
	request := &requests.PlatformRedirectUriRequest{}
	response := &responses.Response{}
	ctx, err := request.Intitiate(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(c, err.Error()))
		return
	}

	if err := request.Validate(ctx); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(c, err.Error()))
		return
	}

	url, err := services.GetOauthUrl(ctx, request)

	c.JSON(http.StatusOK, response.NewSuccessResponse(ctx, "succesfully fetched Oauth Url", url))
}

func Authenticate(c *gin.Context) {
	request := &requests.AuthenticationRequest{}
	response := &responses.Response{}
	ctx, err := request.Intitiate(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(c, err.Error()))
		return
	}

	if err := request.Validate(ctx); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(c, err.Error()))
		return
	}

	data, err := services.Authenticate(ctx, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(ctx, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(ctx, "successfully authenticated", data))
}
