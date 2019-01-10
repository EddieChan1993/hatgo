package link

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	"hatgo/pkg/logs"
	"hatgo/pkg/s"
)

const mysqlLogIH = "[xorm] [info]"

var Db *xorm.Engine

func init() {
	db()
}

func db() {
	var (
		err                                           error
		dbType, dbName, user, pass, host, tablePrefix string
	)

	sec, err := s.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database':%v", err)
	}
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("DB").String()
	user = sec.Key("USER").String()
	pass = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").MustString("")

	connectStr := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		user,
		pass,
		host,
		dbName)
	Db, err = xorm.NewEngine(dbType, connectStr)
	if err != nil {
		log.Printf("%v\n", err)
	}

	log.Printf("%s host:%s user:%s pass:%s db:%s\n", mysqlLogIH, host, user, pass, dbName)
	err = Db.Ping()
	if err != nil {
		log.Printf("%s %v\n", mysqlLogIH, err)
	} else {
		log.Printf("%s %s\n", mysqlLogIH, "mysql's connecting is ok")
	}
	//设置表前缀
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, tablePrefix)
	Db.SetTableMapper(tbMapper)
	logger := xorm.NewSimpleLogger(logs.LogsSql)
	Db.ShowSQL(true)
	Db.SetLogger(logger)
}
