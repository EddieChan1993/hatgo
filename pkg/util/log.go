package util

import (
	"log"
	"os"
)

var (
	InfoC *log.Logger // 重要的信息
	ErrC  *log.Logger // 需要注意的信息
)

func init() {
	//ioutil.Discard不打印
	//os.Stdout输出打印
	InfoC = log.New(os.Stdout,
		"[INFO]",
		log.Ldate|log.Ltime|log.Llongfile)
	ErrC = log.New(os.Stdout,
		"[WARNING]",
		log.Ldate|log.Ltime|log.Llongfile)
}

