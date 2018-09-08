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
	valid := validation.Validation{}
	req := new(ReqTest2)
	c.ShouldBind(req)
	valid.Required(req.Name, "名字").Message("不能为空")
	valid.Range(req.Age, 18, 25, "年龄").Message("不在指定范围")
	valid.Email(req.Email, "邮箱").Message("不合法")
	valid.Mobile(req.Mobile,"电话").Message("不合法")
	valid.IP(req.IP, "IP地址").Message("不合法")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			return fmt.Errorf("%s%s", err.Key, err.Message)
		}
	}
	return nil
}
