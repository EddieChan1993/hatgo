package s

import (
	"time"
	"github.com/go-ini/ini"
	"fmt"
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
	Serverer *Server
)

func init() {
	fmt.Println("--------------------------------------------------------------")
	load()
	validate()
}
