package api

import (
	"english_playground/app/base"
	authentication2 "english_playground/app/controllers/authentication"
	notification2 "english_playground/app/controllers/notification"
	"english_playground/app/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Api struct {
}

func (a Api) Do(router *base.Router) {
	api := router.Group("/api")

	// ping this services
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, struct {
			Status  bool   `json:"status"`
			Message string `json:"message"`
		}{
			Status:  true,
			Message: "pong!",
		})
	})

	// notification service
	notification := router.Group("api/notifications", middlewares.AppMiddleware{}.Do())
	notification.POST("/", notification2.NotificationController{}.Send)
	notification.POST("/subscribe-topic", notification2.NotificationController{}.SubscribeTopic)

	// authentication service
	authentication := router.Group("api/authentications", middlewares.AppMiddleware{}.Do())
	authentication.POST("/token-verification", authentication2.AuthenticationController{}.VerifyToken)

}

func Init() Api {
	return Api{}
}
