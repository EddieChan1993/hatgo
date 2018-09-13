package main_test

import (
	"testing"
	"log"
	"syscall"
	"github.com/fvbock/endless"
	"hatgo/pkg/logging"
	"hatgo/pkg/setting"
	"hatgo/app/routers"
	"hatgo/pkg/link"
	"fmt"
)

var _version_ = "none setting"

func testEndLess(T *testing.T) {
	defer func() {
		link.Db.Close()
		link.Rd.Close()
		logging.LogsReq.Close()
		logging.LogsSql.Close()
		logging.LogsErr.Close()
	}()

	endless.DefaultReadTimeOut = setting.Serverer.ReadTimeout
	endless.DefaultWriteTimeOut = setting.Serverer.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20

	server := endless.NewServer(fmt.Sprintf("%s%s", setting.Serverer.HTTPAdd, setting.Serverer.HTTPPort), routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("HOST is %s", setting.Serverer.HTTPAdd)
		log.Printf("Listening port %s", setting.Serverer.HTTPPort)
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	setting.VersionShow(_version_)
	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("[server stop]%v", err)
	}
}

func testNoEndless(T testing.T) {
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
		log.Fatalf("[server stop]%v", err)
	}
}
