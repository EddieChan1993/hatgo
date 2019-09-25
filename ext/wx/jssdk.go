package wx

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"hatgo/pkg/e"
	"hatgo/pkg/plugin"
	"hatgo/pkg/logs"
	"hatgo/pkg/util"
	"math/rand"
	"strconv"
	"time"
)

//https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421141115
type ResTicket struct {
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

//获取签名和ticket

type ResJSSDK struct {
	Appid     string `json:"appid"`
	Noncestr  string `json:"noncestr"`
	Timestamp int64  `json:"timestamp"`
	Url       string `json:"url"`
	Signature string `json:"signature"`
}

//通过config接口注入权限验证配置
//urlForWeb（当前网页的URL，不包含#及其后面部分）注意地址后面的/，确保js获取的是当前调用的完整地址
func JSSDKConf(urlForWeb string) (*ResJSSDK, error) {
	ticket, err := getTicketForRedis()
	if err != nil {
		return nil, logs.SysErr(err)
	}
	//签名算法
	res := new(ResJSSDK)
	timeInt := time.Now().Unix()
	timeStamp := strconv.FormatInt(timeInt, 10)
	nonceStr := nonceStr()

	res.Signature = Signature(ticket, nonceStr, timeStamp, urlForWeb) //签名
	res.Noncestr = nonceStr
	res.Timestamp = timeInt
	res.Url = urlForWeb
	res.Appid = AppidFlag

	return res, nil
}



//wxpay计算签名的函数
// Signature
func Signature(jsTicket, noncestr, timestamp, url string) string {
	h := sha1.New()
	str := []byte(fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", jsTicket, noncestr, timestamp, url))
	h.Write(str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

//读取缓存获取ticket
func getTicketForRedis() (string, error) {
	v, err := plugin.Rd.Get(e.Ticket).Result()
	if err == redis.Nil {
		ticket, err := getTicket()
		if err != nil {
			return "", logs.SysErr(err)
		}
		err = plugin.Rd.Set(e.Ticket, ticket, time.Second*3000).Err()
		if err != nil {
			return "", logs.SysErr(err)
		}
		return ticket, nil
	} else if err != nil {
		return "", logs.SysErr(err)
	}
	return v, nil
}

//获取ticket
func getTicket() (string, error) {
	ak, err := wx.AccessTokenForComFlag()
	if err != nil {
		return "", logs.SysErr(err)
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi", ak)
	body, err := util.GetCurl(url)
	if err != nil {
		return "", logs.SysErr(err)
	}
	resTicket := new(ResTicket)
	err = json.Unmarshal(body, resTicket)
	if err != nil {
		return "", logs.SysErr(err)
	}

	if resTicket.Errcode != 0 {
		return "", logs.SysErr(err)
	}
	return resTicket.Ticket, nil
}

func RandString(l int) string {
	bs := []byte{}
	for i := 0; i < l; i++ {
		bs = append(bs, chars[rand.Intn(len(chars))])
	}
	return string(bs)
}
