//+build !prod

package setting

import (
	"time"
	"log"
	"github.com/gin-gonic/gin"
)

const DEBUG = gin.DebugMode //调试模式

var RunMode    = DEBUG //运行模式

func LoadServer() {

	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server':%v", err)
	}
	HTTPPort = sec.Key("HTTP_PORT").MustString(":8000")
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

	log.Println("server is running in 【开发模式】")
}
