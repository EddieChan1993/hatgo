package main

import (
	"log"
	"hatgo/models"
	"hatgo/logging"
	"hatgo/routers"
	"hatgo/pkg/setting"
)
var _version_ ="none setting"

func main() {
	defer func() {
		models.Engine.Close()
		logging.Logs.Close()
		logging.SqlLogs.Close()
	}()
	router := routers.InitRouter()
	setting.VersionShow(_version_)
	err := router.Run(setting.HTTPPort)
	if err != nil {
		log.Fatal(err)
	}
}