package e

import "net/http"

//自定义code
const (
	ERROR_AUTH   = http.StatusUnauthorized
	CONNECT_OK   = "connect-ok"
	CONNECT_FAIL_AUTH = "connect-fail-auth"
)

var msgFlags = map[interface{}]string{
	ERROR_AUTH:   "用户认证失败",
	CONNECT_OK:   "连接成功",
	CONNECT_FAIL_AUTH: "非法连接",
}

func GetMsg(code interface{}) string {
	return msgFlags[code]
}
