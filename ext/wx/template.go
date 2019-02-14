package wx

import (
	"encoding/json"
	"fans/pkg/logs"
	"fans/pkg/util"
	"fmt"
	"time"
)

//参考https://developers.weixin.qq.com/miniprogram/dev/api/open-api/template-message/sendTemplateMessage.html
type TempData struct {
	Touser     string      `json:"touser"`      //接收者（用户）的 openid
	TemplateId string      `json:"template_id"` //所需下发的模板消息的id
	Page       string      `json:"page"`        //点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	FormId     string      `json:"form_id"`     //表单提交场景下，为 submit 事件带上的 formId；支付场景下，为本次支付的 prepay_id
	Data       interface{} `json:"data"`        //模板内容，不填则下发空模板
}

//模板响应数据
type ResSendTemp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

//模板体内容
type TempContent struct {
	Value string `json:"value"`
}

//模板消息
//发送模板消息
func sendTemp(openid, fromId, templateId string, data interface{}) error {
	accessToken, err := AccessToken()
	if err != nil {
		return logs.SysErr(err)
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token=%s", accessToken)
	m := new(TempData)
	m.Touser = openid
	m.TemplateId = templateId
	m.FormId = fromId
	m.Data = data

	bytesReq, err := json.Marshal(m)
	if err != nil {
		return logs.SysErr(err)
	}
	//发送unified order请求.统一下单接口
	body,err:=util.PostCurl(url,bytesReq,util.JSONHeader)
	if err != nil {
		return logs.SysErr(err)
	}
	resp2 := new(ResSendTemp)
	err = json.Unmarshal(body, resp2)
	if resp2.Errcode != 0 {
		return logs.SysErr(fmt.Errorf(resp2.Errmsg))
	}
	return nil
}

//发送配送消息模板
//openid 用户openid
//fromId 触发机制id
//paySn 支付单号
//goodsName 配送设备
//addrStr 配送地址
func SendDeliverTemp(openid, fromId, paySn, goodsName, addrStr string) error {
	data := make(map[string]TempContent)
	data["keyword1"] = TempContent{Value: paySn}
	data["keyword2"] = TempContent{Value: goodsName}
	data["keyword3"] = TempContent{Value: util.FormatByStamp(time.Now().Unix()+2*60*60, util.YMD_HIS)}
	data["keyword4"] = TempContent{Value: addrStr}
	return sendTemp(openid, fromId, "3BINmC9N0X06lIn3oBbQacqeEMBaF7ZF_rfTXwpEZCk", data)
}
