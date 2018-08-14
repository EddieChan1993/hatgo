package logging

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"os"
)

type logConfT struct {
	Filename string
	Maxdays  int
	Maxsize  int64
}

var (
	filePath, filePathSql string
	Logs                  *logs.BeeLogger //请求日志
	SqlLogs               *logs.BeeLogger //sql日志
)

type selfLog struct {
	BeeLog *logs.BeeLogger
	File   *os.File
}

func init() {
	//请求日志
	Logs = logs.NewLogger()
	filePath, _ = getLogFilePullPath("req", "app")

	logConf := logConfT{
		Filename: filePath,
		Maxdays:  5,
		Maxsize:  5 * mb,
	}
	b, _ := json.Marshal(logConf)
	//Logs.EnableFuncCallDepth(true) //每行的位置
	Logs.SetLogger(logs.AdapterFile, string(b))
	//Logs.SetLogger(logs.AdapterConsole)
	Logs.Async()

	//sql日志
	SqlLogs = logs.NewLogger()
	filePathSql, _ = getLogFilePullPath("sql", "app")

	logConfSql := logConfT{
		Filename: filePathSql,
		Maxdays:  5,
		Maxsize:  5 * mb,
	}
	bSql, _ := json.Marshal(logConfSql)
	//SqlLogs.EnableFuncCallDepth(true) //每行的位置
	SqlLogs.SetLogger(logs.AdapterFile, string(bSql))
	//SqlLogs.SetLogger(logs.AdapterConsole)
	SqlLogs.Async()
}

//自定义日志文件
func NewSelfLog(logPathName, logFileName string) *selfLog {
	//sql日志
	newLogs := logs.NewLogger()
	filePathSql, file := getLogFilePullPath(logPathName, logFileName)

	logConfSql := logConfT{
		Filename: filePathSql,
		Maxdays:  5,
		Maxsize:  5 * mb,
	}
	bSql, _ := json.Marshal(logConfSql)
	newLogs.EnableFuncCallDepth(true) //每行的位置
	newLogs.SetLogger(logs.AdapterFile, string(bSql))
	newLogs.Async()

	selfLog := &selfLog{
		BeeLog: newLogs,
		File:   file,
	}
	return selfLog
}
