package main_test

import (
	"fmt"
	"github.com/fvbock/endless"
	"hatgo/app/router"
	"hatgo/pkg/plugin"
	"hatgo/pkg/s"
	"hatgo/pkg/logs"
	"log"
	"syscall"
	"testing"
)

const keyVer = "[version]"

var _version_ = "none setting"

func testEndLess(T *testing.T) {
	defer func() {
		plugin.Db.Close()
		plugin.Rd.Close()
		logs.LogsReq.Close()
		logs.LogsSql.Close()
	}()

	endless.DefaultReadTimeOut = s.Service.ReadTimeout
	endless.DefaultWriteTimeOut = s.Service.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20

	log.Printf("%s %s", keyVer, _version_)
	server := endless.NewServer(fmt.Sprintf("%s:%s", s.Service.HTTPAdd, s.Service.HTTPPort), router.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("HOST is %s", s.Service.HTTPAdd)
		log.Printf("Listening port is %s", s.Service.HTTPPort)
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("[server stop]%v", err)
	}
}

func testNoEndless(T testing.T) {
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
