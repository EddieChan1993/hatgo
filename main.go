package main

import (
	"fmt"
	"hatgo/app/router"
	"hatgo/pkg/s"
	"hatgo/pkg/logs"
	"hatgo/pkg/plugin"
	"log"
)

const keyVer = "[version]"

var _version_ = "none setting"

func main() {
	defer func() {
		plugin.Db.Close()
		plugin.Rd.Close()
		logs.LogsReq.Close()
		logs.LogsSql.Close()
	}()

	r := router.InitRouter()
	log.Printf("%s %s", keyVer, _version_)
	err := r.Run(fmt.Sprintf("%s:%s", s.Service.HTTPAdd, s.Service.HTTPPort))
	if err != nil {
		log.Fatalf("[server stop]%v", err)
	}
}