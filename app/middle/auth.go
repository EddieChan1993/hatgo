package middle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hatgo/pkg/e"
)

//Socket验证
func SocketAuth(openId string) bool {
	if openId != "hi" {
		return false
	}
	return true
}


func Auth(c *gin.Context) {
	authCode := http.StatusUnauthorized
	token := c.GetHeader("token")
	if token == "" {
		e.Output(c, authCode, http.StatusText(authCode))
		c.Abort()
		return
	}
	c.Set("uid", 12)
	//c.GetInt64("uid")
	c.Next()
}
