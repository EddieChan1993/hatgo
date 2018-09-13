package setting

import (
	"time"
	"github.com/go-ini/ini"
	"fmt"
	"hatgo/pkg/logging"
)

var (
	err error
	Cfg *ini.File
)

type Server struct {
	HTTPAdd      string
	HTTPPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type QiNiu struct {
	Folder     string
	Host       string
	AccessKey  string
	SecretKey  string
	Bucket     string
	IsUseHttps bool
	ZoneKey    string
}

var (
	QiNiuer  *QiNiu
	Serverer *Server
)

func init() {
	fmt.Println("--------------------------------------------------------------")
	load()
	validate()
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
