package setting

import (
	"time"
	"github.com/go-ini/ini"
	"log"
	"fmt"
	"hatgo/logging"
)

var (
	Cfg          *ini.File
	HTTPPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini", "conf/app.dev.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini':%v", err)
	}
	LoadServer()
}

func VersionShow(v string)  {
	versionLog:= logging.NewSelfLog("version", "app")
	defer func() {
		versionLog.BeeLog.Close()
		versionLog.File.Close()
	}()

	versionStr:=fmt.Sprintf("[version] %s", v)
	versionLog.BeeLog.Info(versionStr)
	fmt.Println(versionStr)
}