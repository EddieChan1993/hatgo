package service

import (
	"github.com/gin-gonic/gin"
	"errors"
	"hatgo/pkg/logging"
)

func GetTestT(c *gin.Context) error {
	//models.AllTest()
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
