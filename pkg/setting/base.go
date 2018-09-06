package setting

import (
	"time"
	"github.com/go-ini/ini"
	"fmt"
	"hatgo/pkg/logging"
)

const HOST = "127.0.0.1" //为空则默认0.0.0.0

var (
	err          error
	Cfg          *ini.File
	HTTPPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
)

func init() {
	fmt.Println("--------------------------------------------------------------")
	LoadServer()
}

func VersionShow(v string) {
	versionLog := logging.NewSelfLog("version", "app")
	defer func() {
		versionLog.BeeLog.Close()
		versionLog.File.Close()
	}()

	versionStr := fmt.Sprintf("[version] %s", v)
	versionLog.BeeLog.Info(versionStr)
	fmt.Println(versionStr)
}
