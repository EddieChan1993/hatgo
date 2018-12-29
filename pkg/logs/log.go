package logs

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"os"
	"runtime"
	"fmt"
	"github.com/astaxie/beego/validation"
	"hatgo/pkg/code"
	"github.com/gin-gonic/gin"
)

//单个日志文件存储，默认256M
type LogConfT struct {
	Filename string
	Maxdays  int
	Level    int
}

var (
	filePath, filePathSql, filePathErr string
	LogsReq                            *logs.BeeLogger //请求日志
	LogsSql                            *logs.BeeLogger //sql日志
	logsErr                            *logs.BeeLogger //err日志
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
	LogsReq = logs.NewLogger()
	filePath, _ = getLogFilePullPath("req", "main")

	logConf := LogConfT{
		Filename: filePath,
		Maxdays:  3,
		Level:    6,
	}
	b, _ := json.Marshal(logConf)
	LogsReq.SetLogger(logs.AdapterFile, string(b))
	LogsReq.Async()
}

//sql日志
func sqlLog() {
	LogsSql = logs.NewLogger()
	filePathSql, _ = getLogFilePullPath("sql", "sql")

	logConfSql := LogConfT{
		Filename: filePathSql,
		Maxdays:  3,
		Level:    6,
	}
	bSql, _ := json.Marshal(logConfSql)
	LogsSql.SetLogger(logs.AdapterFile, string(bSql))
	LogsSql.Async()
}

//err日志
func errLog() {
	logsErr = logs.NewLogger()
	filePathErr, _ = getLogFilePullPath("err", "err")
	logConfErr := LogConfT{
		Filename: filePathErr,
		Maxdays:  3,
		Level:    6,
	}

	bErr, _ := json.Marshal(logConfErr)
	//logsErr.EnableFuncCallDepth(true) //每行的位置
	logsErr.SetLogger(logs.AdapterFile, string(bErr))
	if code.RunMode == gin.DebugMode {
		logConfErrConsole := LogConfT{
			Level: 7,
		}
		bErrC, _ := json.Marshal(logConfErrConsole)
		logsErr.SetLogger(logs.AdapterConsole, string(bErrC))
	}
	logsErr.Async()
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
		Maxdays:  3,
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

//记录err到日志文件，并打印到控制台
func SysErr(err error) error {
	_, file, line, _ := runtime.Caller(1)
	fileLine := fmt.Sprintf("%s:%d", file, line)
	logsErr.Error("%v\n%s", err, fileLine)
	return err
}

//请求验证异常抛出
func ValidErr(errs []*validation.Error) error {
	for _, err := range errs {
		return fmt.Errorf("%s%s", err.Key, err.Message)
	}
	return nil
}
