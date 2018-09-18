package logs

import (
	"fmt"
	"os"
	"log"
)

const mb int64 = 1 << (10 * 2)

var (
	logSavePath = "runtime/logs"
	logFileExt  = "log"
)

func getLogFilePath(logFileName string) string {
	return fmt.Sprintf("%s/%s", logSavePath, logFileName)
}

func getLogFilePullPath(logPathName, logFileName string) (string,*os.File) {
	prefixPath := getLogFilePath(logPathName)
	suffixPath := fmt.Sprintf("%s.%s", logFileName, logFileExt)

	filePath := fmt.Sprintf("%s/%s", prefixPath, suffixPath)
	file:=openLogFile(logPathName, filePath)
	return filePath,file
}

func openLogFile(logPathName, filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir(getLogFilePath(logPathName))
	case os.IsPermission(err):
		log.Fatalf("Permission:%v", err)
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Fail to CreateFile:%v", err)
	}
	return file
}

func mkDir(filePath string) {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+filePath, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
