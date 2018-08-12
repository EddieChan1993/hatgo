package e

const (
	SUCCESS        = 0
	WARNING        = 1
	INVALID_PARAMS = 400

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
)

var msgFlags = map[int]string{
	SUCCESS:                        "ok",
	WARNING:                        "warning",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
}

func GetMsg(code int) string {
	msg, ok := msgFlags[code]
	if ok {
		return msg
	}
	return msgFlags[WARNING]
}
