package service

import (
	"github.com/gin-gonic/gin"
	"hatgo/pkg/e"
	"hatgo/pkg/util"
)

func SHandler(c *gin.Context) {
	wss, token, _ := util.NewWs(c)
	defer func() {
		wss.CloseCoon()
	}()
	uid := token
	channel(wss, uid)
}

func channel(wss *util.Ws, token string) {
	var reqMsg util.Message
	for {
		if err := wss.GetMsg(&reqMsg); err != nil {
			break
		}
		switch reqMsg.Type {
		case "connect":
			wss.BindCoon(token)
			wss.SendSelf(e.GetMsg(e.CONNECT_OK),e.CONNECT_OK)
		default:
			wss.SendSelf("未知操作","none")
		}
	}
}
