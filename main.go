package main

import (
	"english_playground/routes"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	/*
		LOGRUS INIT
	*/
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	println("App started")
	err := godotenv.Load()
	if err != nil {
		println("can't load the env : ", err.Error())
		return
	}
	routes.Init()
}
