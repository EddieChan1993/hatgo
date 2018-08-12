package models

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	"hatgo/pkg/setting"
	"hatgo/logging"
)

var Engine *xorm.Engine

func init() {
	db()
}

func db() {
	var (
		err                                           error
		dbType, dbName, user, pass, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database':%v", err)
	}
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("DB").String()
	user = sec.Key("USER").String()
	pass = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").MustString("")

	connectStr := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8",
		user,
		pass,
		host,
		dbName)
	Engine, err = xorm.NewEngine(dbType, connectStr)
	if err != nil {
		log.Fatal(err)
	}

	err = Engine.Ping()
	if err != nil {
		log.Fatal(err)
	}
	//设置表前缀
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, tablePrefix)
	Engine.SetTableMapper(tbMapper)

	logger :=xorm.NewSimpleLogger(logging.SqlLogs)
	Engine.ShowSQL(true)
	Engine.SetLogger(logger)
}

