package service

import (
	"github.com/gin-gonic/gin"
	"errors"
	"hatgo/pkg/logging"
	"hatgo/app/models"
)

type ReqTest struct {
	One string `json:"one"`
}


func SGetTestT(c *gin.Context) error {
	if 1 == 1 {
		c.Get("uid")
		selfLog := logging.NewSelfLog("test", "cf")
		defer func() {
			selfLog.File.Close()
			selfLog.BeeLog.Close()
		}()
		selfLog.BeeLog.Info("what")
		selfLog.BeeLog.Error("what")
		selfLog.BeeLog.Warning("what")
		return errors.New("hello")
	}
	return nil
}

func FAddTest(c *gin.Context) {
	t := new(models.Test)
	c.ShouldBind(t)
}
