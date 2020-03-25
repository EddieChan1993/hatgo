package middle

import (
	"github.com/gin-gonic/gin"
	"hatgo/pkg/util"
	"net/http"
)

const HTTP_TOKEN = "authToken"//nginx对header头的下划线有过滤配置

func Auth(c *gin.Context) {
	authCode := http.StatusUnauthorized
	token := c.GetHeader(HTTP_TOKEN)
	if token == "" {
		util.Output(c, authCode, http.StatusText(authCode))
		c.Abort()
		return
	}
	c.Set("uid", 12)
	//c.GetInt64("uid")
	c.Next()
}
