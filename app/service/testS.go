package service

import (
	"github.com/gin-gonic/gin"
	"hatgo/pkg/logging"
	"fmt"
	"github.com/astaxie/beego/validation"
)

type ReqTest struct {
	One string `json:"one"`
}

type ReqTest2 struct {
	Name   string `json:"name"`
	Age    int    `"json:"age"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
	IP     string `json:"ip"`
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
		return fmt.Errorf("what")
	}
	return nil
}

func FAddTest(c *gin.Context) error {
	v:=new(validation.Validation)
	req := new(ReqTest2)
	c.ShouldBind(req)
	v.Required(req.Name, "名字")
	v.Range(req.Age, 18, 25, "年龄")
	v.Email(req.Email, "")
	v.Mobile(req.Mobile,"")
	v.IP(req.IP, "")
	if v.HasErrors() {
		for _, err := range v.Errors {
			return fmt.Errorf("%s%s", err.Key, err.Message)
		}
	}
	return nil
}
