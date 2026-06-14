package routers

import (
	"auth-service/handler"
	"auth-service/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine) {
	engine.Use(middleware.SetContext())

	oauth := engine.Group("/oauth")
	{
		oauth.GET("/url", handler.GetOauthUrl)
		oauth.GET("/authenticate", handler.Authenticate)
	}
}
