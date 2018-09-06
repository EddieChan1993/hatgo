/**
	常用函数
 */
package util

import (
	"log"
	"strconv"
)

//string->float64
func StringToFloat64(str string) float64 {
	floatVal,err:=strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatal(err)
	}
	return floatVal
}

//string->int64
func StringToInt64(str string) int64 {
	intVal, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return intVal
}

//float64->int64
func Float64ToInt64(floatParam float64) int64 {
	str:=strconv.FormatFloat(floatParam, 'f', -1, 64)
	return StringToInt64(str)
}

//int->float64
func IntToFloat64(intParam int)float64 {
	i:=int64(intParam)
	str:=strconv.FormatInt(i, 10)
	return StringToFloat64(str)
}

//int64->float64
func Int64ToFloat64(intParam int64) float64 {
	str:=strconv.FormatInt(intParam, 10)
	return StringToFloat64(str)
}

//byte->string
func ByteString(p []byte) string {
	//如果没有空字节，直接使用string(p)
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}