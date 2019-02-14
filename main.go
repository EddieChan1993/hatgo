package main

import (
	"fmt"
	"hatgo/app/router"
	"hatgo/pkg/s"
	"hatgo/pkg/link"
	"hatgo/pkg/logs"
	"log"
)

const keyVer = "[version]"

var _version_ = "none setting"

func main() {

	defer func() {
		link.Db.Close()
		link.Rd.Close()
		logs.LogsReq.Close()
		logs.LogsSql.Close()
	}()

	router := router.InitRouter()
	log.Printf("%s %s", keyVer, _version_)
	err := router.Run(fmt.Sprintf("%s:%s", s.Serverer.HTTPAdd, s.Serverer.HTTPPort))
	if err != nil {
		log.Fatalf("[server stop]%v", err)
	}
}
