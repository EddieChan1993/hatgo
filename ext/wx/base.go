package wx

import (
	"encoding/json"
	"fans/pkg/logs"
	"fans/pkg/util"
	"fmt"
)

//获取openid
//openid的响应数据
type ResOpenId struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Openid  string `json:"openId"`
}

//accessToken
type ResAccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"exprires_in"` //过期时间
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
}

func OpenidXCX(code string) (string, error) {
	return authOpenid(code, appidXCX)
}

//获取openid
func authOpenid(code, appid string) (string, error) {
	host := "https://api.weixin.qq.com/sns/jscode2session"
	formUrl := "%s?appid=%s&secret=%s&js_code=%s&grant_type=uthorization_code"
	url := fmt.Sprintf(formUrl, host, appid, appSecretXCX, code)
	resOpenid := new(ResOpenId)
	body, err := util.GetCurl(url)
	err = json.Unmarshal(body, resOpenid)
	if err != nil {
		return "", logs.SysErr(err)
	}
	if resOpenid.Errcode != 0 {
		return "", logs.SysErr(fmt.Errorf(resOpenid.Errmsg))
	}
	return resOpenid.Openid, nil
}

//获取access_token
func AccessToken() (string, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appidXCX, appSecretXCX)
	body, err := util.GetCurl(url)
	respM := new(ResAccessToken)
	err = json.Unmarshal(body, respM)
	if err != nil {
		return "", logs.SysErr(err)
	}
	if respM.Errcode != 0 {
		return "", logs.SysErr(fmt.Errorf(respM.Errmsg))
	}
	return respM.AccessToken, nil
}
