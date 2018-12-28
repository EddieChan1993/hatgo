package e

import "net/http"

//自定义code
const (
	ErrorAuth = http.StatusUnauthorized
)

var msgFlags = map[int]string{
	ErrorAuth: "用户认证失败",
}

func GetMsg(code int) string {
	return msgFlags[code]
}
