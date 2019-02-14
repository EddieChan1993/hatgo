package main_test

import (
	"fmt"
	"github.com/fvbock/endless"
	"hatgo/app/router"
	"hatgo/pkg/s"
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

	endless.DefaultReadTimeOut = s.Serverer.ReadTimeout
	endless.DefaultWriteTimeOut = s.Serverer.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20

	log.Printf("%s %s", keyVer, _version_)
	server := endless.NewServer(fmt.Sprintf("%s:%s", s.Serverer.HTTPAdd, s.Serverer.HTTPPort), router.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("HOST is %s", s.Serverer.HTTPAdd)
		log.Printf("Listening port is %s", s.Serverer.HTTPPort)
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
	err := router.Run(fmt.Sprintf("%s:%s", s.Serverer.HTTPAdd, s.Serverer.HTTPPort))
	if err != nil {
		log.Fatalf("[server stop]%v", err)
	}
}
