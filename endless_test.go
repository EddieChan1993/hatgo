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
		logging.Logs.Close()
		logging.SqlLogs.Close()
	}()

	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20

	server := endless.NewServer(fmt.Sprintf("%s%s", setting.HOST, setting.HTTPPort), routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Println("server is running in", setting.RunMode)
		log.Println("Listening port", setting.HTTPPort)
		log.Println("Actual pid is", syscall.Getpid())
	}
	setting.VersionShow(_version_)
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Server err:%v", err)
	}
}

func testNoEndless(T testing.T) {
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
