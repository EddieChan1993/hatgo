package main_test

import (
	"testing"
	"log"
	"syscall"
	"github.com/fvbock/endless"
	"hatgo/pkg/logs"
	"hatgo/pkg/conf"
	"hatgo/app/router"
	"hatgo/pkg/link"
	"fmt"
)
const keyVer = "[version]"
var _version_ = "none setting"

func testEndLess(T *testing.T) {
	defer func() {
		link.Db.Close()
		link.Rd.Close()
		logs.LogsReq.Close()
		logs.LogsSql.Close()
	}()

	endless.DefaultReadTimeOut = conf.Serverer.ReadTimeout
	endless.DefaultWriteTimeOut = conf.Serverer.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20

	log.Printf("%s %s",keyVer,_version_)
	server := endless.NewServer(fmt.Sprintf("%s%s", conf.Serverer.HTTPAdd, conf.Serverer.HTTPPort), router.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("HOST is %s", conf.Serverer.HTTPAdd)
		log.Printf("Listening port %s", conf.Serverer.HTTPPort)
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("[server stop]%v", err)
	}
}

func testNoEndless(T testing.T) {
	defer func() {
		link.Db.Close()
		link.Rd.Close()
		logs.LogsReq.Close()
		logs.LogsSql.Close()
	}()

	router := router.InitRouter()
	log.Printf("%s %s",keyVer,_version_)
	err := router.Run(fmt.Sprintf("%s%s", conf.Serverer.HTTPAdd, conf.Serverer.HTTPPort))
	if err != nil {
		log.Fatalf("[server stop]%v", err)
	}
}
