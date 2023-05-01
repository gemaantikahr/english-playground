package middlewares

import (
	"english_playground/app/response"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type AppMiddleware struct {
}

func (m AppMiddleware) Do() gin.HandlerFunc {

	return func(c *gin.Context) {

		appKey := c.GetHeader("app-key")

		if os.Getenv("APP_KEY") != appKey || appKey == "" {

			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Api().Error("unauthorized", nil))
			return

		}

	}
}
