package main_test

import (
	"fmt"
	"github.com/fvbock/endless"
	"hatgo/app/router"
	"hatgo/pkg/c"
	"hatgo/pkg/link"
	"hatgo/pkg/logs"
	"log"
	"syscall"
	"testing"
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

	endless.DefaultReadTimeOut = c.Serverer.ReadTimeout
	endless.DefaultWriteTimeOut = c.Serverer.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20

	log.Printf("%s %s", keyVer, _version_)
	server := endless.NewServer(fmt.Sprintf("%s:%s", c.Serverer.HTTPAdd, c.Serverer.HTTPPort), router.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("HOST is %s", c.Serverer.HTTPAdd)
		log.Printf("Listening port %s", c.Serverer.HTTPPort)
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
	log.Printf("%s %s", keyVer, _version_)
	err := router.Run(fmt.Sprintf("%s:%s", c.Serverer.HTTPAdd, c.Serverer.HTTPPort))
	if err != nil {
		log.Fatalf("[server stop]%v", err)
	}
}
