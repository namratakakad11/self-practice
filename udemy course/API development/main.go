package main

import (
	"APIDEVELOPMENT/app"
	"APIDEVELOPMENT/logger"
)

func main() {

	logger.Info("starting application")
	app.Start()

}
