package service

import (
	"github.com/gin-gonic/gin"
	"hatgo/pkg/util"
	"hatgo/pkg/logs"
	"fmt"
)

func SHandler(c *gin.Context) {
	wss := util.NewWs(c)
	uid := util.CoonId()

	defer func() {
		logs.SysErr(fmt.Errorf("%s离开", uid))
		wss.CloseUid(uid)
	}()
	var msg util.Message
	for {
		if err := wss.GetMsg(&msg); err != nil {
			logs.SysErr(err)
			break
		}
		fmt.Println(msg)
		switch msg.Type {
		case "connect":
			wss.BindUid(uid)
			msg = util.Message{
				Content: "连接成功",
				Type:    "connect-ok",
			}
			wss.SendToUid(uid, &msg)
		}
	}
}
