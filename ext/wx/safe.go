package wx

import (
	"encoding/json"
	"fmt"
	"hatgo/pkg/logs"
	"hatgo/pkg/util"
	"mime/multipart"
)

/**
	内容安全接口文档
	https://mp.weixin.qq.com/cgi-bin/announce?token=233192696&action=getannouncement&key=11522142966rk3L2&version=1&lang=zh_CN&platform=2
 */

const (
	msgWarning = "内容中存在敏感词，无法提交"
	imgWarning = "图片较为敏感，无法提交"
)

//内容检测请求
type ReqMsgCheck struct {
	Content string `json:"content"`
}

//图片检测请求
type ReqImgCheck struct {
	Media *multipart.FileHeader `json:"media"`
}

//模板响应数据
type ResData struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

//内容敏感词检测
func MsgCheck(content []byte) error {
	ak, err := AccessToken()
	if err != nil {
		return logs.SysErr(err)
	}
	req := new(ReqMsgCheck)
	req.Content = string(content)
	url := fmt.Sprintf("https://api.weixin.qq.com/wxa/msg_sec_check?access_token=%s", ak)
	bt, err := json.Marshal(req)
	if err != nil {
		return logs.SysErr(err)
	}
	reqParams := new(util.ReqParams)
	reqParams.Url = url
	reqParams.Method = util.POST
	reqParams.Header = util.JSONHeader
	reqParams.Params = bt
	reqObj, err := reqParams.InitRequest()
	if err != nil {
		return logs.SysErr(err)
	}
	body, err := reqObj.Do()
	if err != nil {
		return logs.SysErr(err)
	}
	res := new(ResData)
	err = json.Unmarshal(body, res)
	if err != nil {
		return logs.SysErr(err)
	}
	if res.Errcode == 87014 {
		return logs.SysErr(fmt.Errorf(msgWarning))
	}else if res.Errcode!=0 {
		return logs.SysErr(fmt.Errorf(res.Errmsg))
	}
	return nil
}

