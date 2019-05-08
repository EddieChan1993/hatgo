package wx

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"hatgo/pkg/e"
	"hatgo/pkg/link"
	"hatgo/pkg/logs"
	"hatgo/pkg/util"
	"sort"
	"strconv"
	"strings"
	"time"
)

//jssdk相关的准备工作

//https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421141115
type ResTicket struct {
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

//获取签名和ticket
type ResJSSDK struct {
	Signature string `json:"signature"`
	Timestamp int    `json:"timestamp"`
	Noncestr  string `json:"noncestr"`
}

//通过config接口注入权限验证配置
//urlForWeb（当前网页的URL，不包含#及其后面部分）
func JSSDKConf(urlForWeb string) (*ResJSSDK, error) {
	ticket, err := getTicketForRedis()
	if err != nil {
		return nil, logs.SysErr(err)
	}
	//签名算法
	resMap := make(map[string]interface{}, 0)
	resMap["jsapi_ticket"] = ticket
	resMap["nonceStr"] = nonceStr()
	resMap["timeStamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	resMap["url"] = urlForWeb

	res := new(ResJSSDK)
	res.Signature = wxCalcSign(resMap) //签名
	res.Noncestr = nonceStr()
	res.Timestamp = resMap["timeStamp"].(int)
	return res, nil
}
//wxpay计算签名的函数
func wxCalcSign(mReq map[string]interface{}) (sign string) {
	//STEP 1, 对key进行升序排序.
	sortedKeys := make([]string, 0)
	for k := range mReq {
		sortedKeys = append(sortedKeys, k)
	}

	sort.Strings(sortedKeys)

	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sortedKeys {
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	//STEP4, 进行sha1签名并且将所有字符转为大写.
	md5Ctx := sha1.New()
	md5Ctx.Write([]byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	return upperSign
}

//读取缓存获取ticket
func getTicketForRedis() (string, error) {
	v, err := link.Rd.Get(e.Ticket).Result()
	if err == redis.Nil {
		ticket, err := getTicket()
		if err != nil {
			return "", logs.SysErr(err)
		}
		err = link.Rd.Set(e.Ticket, ticket, 4500*time.Second).Err()
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
	ak, err := AccessTokenForComFlag()
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
