package main

import (
	"hatgo/pkg/link"
	"fmt"
	"hatgo/pkg/setting"
	"log"
	"hatgo/pkg/logging"
	"hatgo/app/routers"
	"os"
)

var _version_ = "none setting"

func main() {

	defer func() {
		link.Db.Close()
		link.Rd.Close()
		logging.LogsReq.Close()
		logging.LogsSql.Close()
		logging.LogsErr.Close()
	}()

	router := routers.InitRouter()
	setting.VersionShow(_version_)
	err := router.Run(fmt.Sprintf("%s%s", setting.Serverer.HTTPAdd, setting.Serverer.HTTPPort))
	if err != nil {
		log.Printf("%v\n", err)
		os.Exit(1)
	}
}
