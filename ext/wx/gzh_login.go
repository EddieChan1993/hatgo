package wx

/**
公众号相关
https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140842
1 第一步：用户同意授权，获取code

2 第二步：通过code换取网页授权access_token

3 第三步：刷新access_token（如果需要）

4 第四步：拉取用户信息(需scope为 snsapi_userinfo)

5 附：检验授权凭证（access_token）是否有效
 */
import (
	"encoding/json"
	"fmt"
	"hatgo/pkg/logs"
	"hatgo/pkg/util"
	url2 "net/url"
)

//accessToken接口返回数据
type RespATByFlag struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Openid      string `json:"openid"`
	Errcode     int    `json:"errocode"`
	Errmsg      string `json:"errmsg"`
}

//用户基本信息
type RespUinfoByFlag struct {
	Openid     string `json:"openid"`
	Nickname   string `json:"nickname"`
	Sex        string `json:"sex"` //0-未知 1-男 2-女
	Province   string `json:"province"`
	Headimgurl string `json:"headimgurl"`
	Errcode    int    `json:"errocode"`
	Errmsg     string `json:"errmsg"`
	Unionid    string `json:"unionid"`
}

//获取code
func GetCode(redirectUrl string) string {
	api := "https://open.weixin.qq.com/connect/oauth2/authorize"
	params := url2.Values{}
	params.Add("redirect_uri", redirectUrl)
	p := params.Encode()
	url := fmt.Sprintf("%s?appid=%s&%s&response_type=code&scope=SCOPE&state=STATE#wechat_redirect", api, wx.AppidFlag, p)
	return url
}

//获取accessToken
func GetAccessToken(code string) (*RespATByFlag, error) {
	api := "https://api.weixin.qq.com/sns/oauth2/access_token"
	url := fmt.Sprintf("%s?appid=%s&secret=%s&code=%s&grant_type=authorization_code", api, wx.AppidFlag, wx.AppSecretFlag, code)
	body, err := util.GetCurl(url)
	if err != nil {
		return nil, logs.SysErr(err)
	}
	data := new(RespATByFlag)
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, logs.SysErr(err)
	}
	if data.Errcode != 0 {
		return nil, logs.SysErr(fmt.Errorf(data.Errmsg))
	}
	return data, nil
}

//获取用户基本信息
func GetUinfoByFlag(ak string) (*RespUinfoByFlag, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=OPENID&lang=zh_CN", ak)
	body, err := util.GetCurl(url)
	if err != nil {
		return nil, logs.SysErr(err)
	}
	data := new(RespUinfoByFlag)
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, logs.SysErr(err)
	}
	if data.Errcode != 0 {
		return nil, logs.SysErr(fmt.Errorf(data.Errmsg))
	}
	return data, nil
}
