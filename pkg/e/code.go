package e

import "net/http"

//自定义code
const (
	ERROR_AUTH                     = http.StatusUnauthorized
)

var msgFlags = map[int]string{
	ERROR_AUTH:                     "用户认证失败",
}

func GetMsg(code int) string {
	msg, ok := msgFlags[code]
	if ok {
		return msg
	}
	return msgFlags[code]
}
