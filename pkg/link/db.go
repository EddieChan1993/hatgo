package link

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	"hatgo/pkg/setting"
	"hatgo/pkg/logging"
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
	Db, err = xorm.NewEngine(dbType, connectStr)
	if err != nil {
		log.Fatal(err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}else{
		fmt.Printf("%s %s\n",mysqlLogIH,"mysql's connecting is ok")
	}
	//设置表前缀
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, tablePrefix)
	Db.SetTableMapper(tbMapper)

	logger :=xorm.NewSimpleLogger(logging.SqlLogs)
	Db.ShowSQL(true)
	Db.SetLogger(logger)
}
