package weixin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"foleng/pkg/logs"
	"io/ioutil"
	"net/http"
)

//参考https://developers.weixin.qq.com/miniprogram/dev/api/open-api/template-message/sendTemplateMessage.html
type TempData struct {
	Touser     string      `json:"touser"`      //接收者（用户）的 openid
	TemplateId string      `json:"template_id"` //所需下发的模板消息的id
	Page       string      `json:"page"`        //点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	FormId     string      `json:"form_id"`     //表单提交场景下，为 submit 事件带上的 formId；支付场景下，为本次支付的 prepay_id
	Data       interface{} `json:"data"`        //模板内容，不填则下发空模板
}

//模板消息
//发送模板消息
func SentTemp(accessToken string) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token=%s", accessToken)
	m := new(TempData)
	m.Touser = "o3zap5V5LsoY4hSkYI9ZMCidX-8I"
	m.TemplateId = "3BINmC9N0X06lIn3oBbQacqeEMBaF7ZF_rfTXwpEZCk"
	m.FormId = "wx11193222063664ab1b1cba600983195080"

	bytesReq, err := json.Marshal(m)
	if err != nil {
		logs.SysErr(err)
	}
	fmt.Println(string(bytesReq))
	//发送unified order请求.统一下单接口
	req, err := http.NewRequest("POST", url, bytes.NewReader(bytesReq))
	if err != nil {
		logs.SysErr(err)
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	c := http.Client{}
	resp, err := c.Do(req)
	defer resp.Body.Close()
	if err != nil {
		logs.SysErr(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.SysErr(err)
	}
	fmt.Println(string(body))
}
