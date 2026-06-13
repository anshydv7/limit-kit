package middleware

import (
	"auth-service/utils"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func SetContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := utils.GenerateRandomId(32)
		ctx := utils.ContextWithRequestId(c.Request.Context(), reqID)
		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		c.Set("context", ctx)
		c.Set("request_id", reqID)
		c.Next()
	}
}
