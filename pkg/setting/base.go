package setting

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