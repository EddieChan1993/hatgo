package service

import (
	"github.com/gin-gonic/gin"
	"hatgo/pkg/e"
	"hatgo/pkg/util"
	"hatgo/app/middle"
)

func SHandler(c *gin.Context) {
	wss, token, _ := util.NewWs(c)
	defer func() {
		wss.CloseCoon()
	}()
	b := middle.SocketAuth(token)
	if !b {
		resMsg := &util.Message{
			Content: e.GetMsg(e.CONNECT_FAIL_AUTH),
			Type:    e.CONNECT_FAIL_AUTH,
		}
		wss.SendSelf(resMsg)
		wss.CloseCoon()
	}
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
			resMsg := &util.Message{
				Content: e.GetMsg(e.CONNECT_OK),
				Type:    e.CONNECT_OK,
			}
			wss.SendSelf(resMsg)
		default:
			resMsg := &util.Message{
				Content: "未知操作",
				Type:    "none",
			}
			wss.SendSelf(resMsg)
		}
	}
}
