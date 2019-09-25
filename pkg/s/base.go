package s

import (
	"fmt"
	"github.com/go-ini/ini"
	"time"
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

var (
	Service *Server
)

func init() {
	fmt.Println("--------------------------------------------------------------")
	load()
	//plugin.Validate()
}
