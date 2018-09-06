package service

import (
	"github.com/gin-gonic/gin"
	"errors"
	"hatgo/pkg/logging"
)

type TestR struct {
	One string `json:"one"`
}

func GetTestT(c *gin.Context) error {
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

func AddTest(c *gin.Context) {
	t := new(TestR)
	c.ShouldBind(t)
}
