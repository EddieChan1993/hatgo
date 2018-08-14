//+build !prod

package setting

import (
	"time"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
)

const RunMode = gin.DebugMode //调试模式

func LoadServer() {
	Cfg, err = ini.Load("conf/app.ini", "conf/app.dev.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini':%v", err)
	}
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server':%v", err)
	}
	HTTPPort = sec.Key("HTTP_PORT").MustString(":8000")
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

	log.Println("server is running in 【开发模式】")
}
