package middle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hatgo/pkg/e"
	"hatgo/pkg/util"
)

//Socket验证
func SocketAuth(wss *util.Ws, openId string){
	if openId != "hi" {
		wss.SendSelf(e.GetMsg(e.CONNECT_FAIL_AUTH),e.CONNECT_FAIL_AUTH)
		wss.CloseCoon()
	}
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
