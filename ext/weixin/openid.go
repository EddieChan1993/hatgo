package weixin

import (
	"fmt"
	"encoding/json"
	"foleng/pkg/util"
	"foleng/pkg/logs"
)

//获取openid
//openid的响应数据
type ResOpenId struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Openid  string `json:"openId"`
}

//获取openid
func AuthOpenid(code string) (string, error) {
	var d []byte
	host := "https://api.weixin.qq.com/sns/jscode2session"
	formUrl := "%s?appid=%s&secret=%s&js_code=%s&grant_type=uthorization_code"
	url := fmt.Sprintf(formUrl, host, appid, appSecret, code)
	resOpenid := new(ResOpenId)
	d, err := util.HttpCurl(url).Get()
	if err != nil {
		return "", logs.SysErr(err)
	}
	err = json.Unmarshal(d, resOpenid)
	if err != nil {
		return "", logs.SysErr(err)
	}
	if resOpenid.Errcode != 0 {
		return "", logs.SysErr(fmt.Errorf(resOpenid.Errmsg))
	}
	return resOpenid.Openid, nil
}
