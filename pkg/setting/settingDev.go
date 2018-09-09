//+build !prod

package setting

import (
	"log"
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
	"github.com/go-ini/ini"
)

const RunMode = gin.DebugMode //调试模式
func load() {
	Cfg, err = ini.Load("conf/app.ini", "conf/app.dev.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini':%v", err)
	}
	loadServer()
	loadQiniu()
}
func loadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server':%v", err)
	}
	HTTPAdd = sec.Key("HTTP_ADDR").MustString("")
	HTTPPort = sec.Key("HTTP_PORT").MustString(":8000")
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
	fmt.Println("server is running in 【开发模式】")
}

func loadQiniu() {
	sec, err := Cfg.GetSection("qiniu")
	if err != nil {
		log.Fatalf("Fail to get section 'qiniu':%v", err)
	}
	Host = sec.Key("host").MustString("")
	AccessKey = sec.Key("accessKey").MustString("")
	SecretKey = sec.Key("secretKey").MustString("")
	Bucket = sec.Key("bucket").MustString("")
	Folder = sec.Key("folder").MustString("")
	IsUseHttps = sec.Key("host").MustBool(false)
	ZoneKey = sec.Key("zone").MustString("")
}
