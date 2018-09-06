package logging

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"os"
)

type LogConfT struct {
	Filename string
	Maxdays  int
	Maxsize  int64
	Level    int
}

var (
	filePath, filePathSql, filePathErr string
	Logs                               *logs.BeeLogger //请求日志
	SqlLogs                            *logs.BeeLogger //sql日志
	ErrLogs                            *logs.BeeLogger //err日志
)

type selfLog struct {
	BeeLog *logs.BeeLogger
	File   *os.File
}

func init() {
	reqLog()
	sqlLog()
	errLog()
}

//请求日志
func reqLog() {
	Logs = logs.NewLogger()
	filePath, _ = getLogFilePullPath("req", "req")

	logConf := LogConfT{
		Filename: filePath,
		Maxdays:  5,
		Maxsize:  5 * mb,
		Level:6,
	}
	b, _ := json.Marshal(logConf)
	Logs.SetLogger(logs.AdapterFile, string(b))
	Logs.Async()
}
//sql日志
func sqlLog() {
	SqlLogs = logs.NewLogger()
	filePathSql, _ = getLogFilePullPath("sql", "sql")

	logConfSql := LogConfT{
		Filename: filePathSql,
		Maxdays:  5,
		Maxsize:  5 * mb,
		Level:6,
	}
	bSql, _ := json.Marshal(logConfSql)
	SqlLogs.SetLogger(logs.AdapterFile, string(bSql))
	SqlLogs.Async()
}

//err日志
func errLog()  {
	ErrLogs = logs.NewLogger()
	filePathErr, _ = getLogFilePullPath("err", "err")
	logConfErr := LogConfT{
		Filename: filePathErr,
		Maxdays:  5,
		Maxsize:  5 * mb,
		Level:6,
	}
	logConfErrConsole := LogConfT{
		Level: 7,
	}
	bErr, _ := json.Marshal(logConfErr)
	bErrC, _ := json.Marshal(logConfErrConsole)
	ErrLogs.EnableFuncCallDepth(true) //每行的位置
	ErrLogs.SetLogger(logs.AdapterConsole, string(bErrC))
	ErrLogs.SetLogger(logs.AdapterFile, string(bErr))
}
/**
 	log.Emergency("Emergency")
	log.Alert("Alert")
	log.Critical("Critical")
	log.Error("Error")
	log.Warning("Warning")
	log.Notice("Notice")
	log.Informational("Informational")
	log.Debug("Debug")
 */

//自定义日志文件
func NewSelfLog(logPathName, logFileName string) *selfLog {
	//sql日志
	newLogs := logs.NewLogger()
	filePathSql, file := getLogFilePullPath(logPathName, logFileName)

	logConf := LogConfT{
		Filename: filePathSql,
		Maxdays:  5,
		Maxsize:  5 * mb,
	}
	b, _ := json.Marshal(logConf)
	newLogs.EnableFuncCallDepth(true) //每行的位置
	newLogs.SetLogger(logs.AdapterFile, string(b))
	newLogs.Async()

	selfLog := &selfLog{
		BeeLog: newLogs,
		File:   file,
	}
	return selfLog
}
