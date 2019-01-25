package model

import (
	"fmt"
	"time"
)

//"Error 1062: Duplicate entry 'fans-adm2in1' for key 'username'",
const SqlCodeErr = "Error 1062"

type XormDate time.Time

func (this XormDate) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

//是否是sql字段值重复错误
func IsDuplicateError(err error) bool {
	bt := []byte(err.Error())
	if string(bt[:10]) == SqlCodeErr {
		return true
	}
	return false
}
