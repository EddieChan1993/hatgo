package main

import (
	"hatgo/pkg/link"
	"fmt"
	"hatgo/pkg/setting"
	"log"
	"hatgo/pkg/logs"
	"hatgo/app/routers"
)

var _version_ = "none setting"

func main() {

	defer func() {
		link.Db.Close()
		link.Rd.Close()
		logs.LogsReq.Close()
		logs.LogsSql.Close()
	}()

	router := routers.InitRouter()
	err := router.Run(fmt.Sprintf("%s%s", setting.Serverer.HTTPAdd, setting.Serverer.HTTPPort))
	if err != nil {
		log.Fatalf("[server stop]%v", err)
	}
}
