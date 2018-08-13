package main

import (
	"log"
	"hatgo/logging"
	"hatgo/app/routers"
	"hatgo/pkg/setting"
	"hatgo/pkg/link"
)
var _version_ ="none setting"

func main() {
	defer func() {
		link.Db.Close()
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