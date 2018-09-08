package setting

import (
	"time"
	"github.com/go-ini/ini"
	"fmt"
	"hatgo/pkg/logging"
)

var (
	err          error
	Cfg          *ini.File
	HTTPAdd      string
	HTTPPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
)

func init() {
	fmt.Println("--------------------------------------------------------------")
	loadServer()
	initValidate()
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
