package e

import "github.com/gin-gonic/gin"

func ResSuccess(data interface{}) map[string]interface{} {
	return gin.H{
		"code": SUCCESS,
		"msg":  GetMsg(SUCCESS),
		"data": data,
	}
}

func ResWarning(data interface{}) map[string]interface{} {
	return gin.H{
		"code": WARNING,
		"msg":  GetMsg(WARNING),
		"data": data,
	}
}

func ResOutput(code int, data interface{}) map[string]interface{} {
	return gin.H{
		"code": code,
		"msg":  GetMsg(code),
		"data": data,
	}
}
