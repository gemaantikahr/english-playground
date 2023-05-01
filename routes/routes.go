package routes

import (
	"english_playground/app/base"
	"english_playground/routes/api"
	"english_playground/routes/web"
	"os"
)

type RouteInterface interface {
	Do(router *base.Router)
}

func Init() {
	router := base.NewRouter()
	api.Init().Do(router)
	web.Init().Do(router)

	router.Run(":" + os.Getenv("APP_PORT"))
}
