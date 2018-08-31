package main

import (
	"log"
	"hatgo/pkg/logging"
	"hatgo/app/routers"
	"hatgo/pkg/setting"
	"hatgo/pkg/link"
	"fmt"
)

var _version_ = "none setting"

func main() {

	defer func() {
		link.Db.Close()
		link.Rd.Close()
		logging.Logs.Close()
		logging.SqlLogs.Close()
	}()

	router := routers.InitRouter()
	setting.VersionShow(_version_)
	err := router.Run(fmt.Sprintf("%s%s", setting.HOST, setting.HTTPPort))
	if err != nil {
		log.Fatal(err)
	}
}
