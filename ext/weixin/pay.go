package weixin

import (
	"fmt"
	"crypto/md5"
	"strings"
	"encoding/hex"
	"sort"
	"time"
	"encoding/xml"
	"net/http"
	"io/ioutil"
	"bytes"
	"hatgo/pkg/util"
	"hatgo/pkg/logs"
	"strconv"
)

const SUCCESS = "SUCCESS"

//首先定义一个UnifyOrderReq用于填入我们要传入的参数。
type UnifyOrderReq struct {
	Appid          string `xml:"appid"`
	Openid         string `xml:"openid"`
	Body           string `xml:"body"`
	MchId          string `xml:"mch_id"`
	NonceStr       string `xml:"nonce_str"`
	NotifyUrl      string `xml:"notify_url"`
	TradeType      string `xml:"trade_type"`
	SpbillCreateIp string `xml:"spbill_create_ip"`
	TotalFee       int    `xml:"total_fee"`
	OutTradeNo     string `xml:"out_trade_no"`
	Sign           string `xml:"sign"`
}

//订单商品
type WxOrderGoods struct {
	Body           string //商品名
	TotalFee       int    //支付价格,单位分
	TradeNo        string //订单号
	SpbillCreateIp string //设备ip
	NotifyUrl      string //支付回调 eg:"https://www.yourserver.com/wxpayNotify"
}

//统一下单接口返回数据
type ReturnData struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	NonceStr   string `xml:"nonce_str"` //随机字符串
	PrepayId   string `xml:"prepay_id"` //预支付交易会话标识
}

//客户端支付所需返回数据
type ResPayData struct {
	TimeStamp string `json:"time_stamp"` //时间戳
	NonceStr  string `json:"nonce_str"`  //随机数据
	PrepayId  string `json:"prepay_id"`  //统一下单接口返回的 prepay_id 参数值
	PaySign   string `json:"pay_sign"`   //签名
}

//支付成功回调结构数据
type NoticeData struct {
	ReturnCode    string `xml:"return_code"` //交易成功返回SUCCESS
	ReturnMsg     string `xml:"return_msg"`
	OutTradeNo    string `xml:"out_trade_no"`   //商户订单号，对应自定义订单号TradeNo
	TransactionId string `xml:"transaction_id"` //微信交易单号
	TotalFee      int    `xml:"total_fee"`      //订单金额,单位：分
	Sign          string `xml:"sign"`           //签名
}

//小程序支付
func PayXCX(openId string, orderGoods *WxOrderGoods) (*ResPayData, error) {
	return unifiedOrder(openId, appidXCX, jsapi, orderGoods)
}

//统一下单
func unifiedOrder(openId, appid, tradeType string, orderGoods *WxOrderGoods) (*ResPayData, error) {
	var data UnifyOrderReq
	//统一下单请求参数
	data.Appid = appid
	data.MchId = mchId
	data.Openid = openId
	data.NonceStr = nonceStr()
	data.TradeType = tradeType
	data.Body = orderGoods.Body
	data.NotifyUrl = orderGoods.NotifyUrl
	data.SpbillCreateIp = orderGoods.SpbillCreateIp
	data.TotalFee = orderGoods.TotalFee //单位是分，这里是1毛钱
	data.OutTradeNo = orderGoods.TradeNo

	//签名算法
	m := make(map[string]interface{}, 0)
	m["appid"] = data.Appid
	m["body"] = data.Body
	m["mch_id"] = data.MchId
	m["openid"] = data.Openid
	m["notify_url"] = data.NotifyUrl
	m["trade_type"] = data.TradeType
	m["spbill_create_ip"] = data.SpbillCreateIp
	m["total_fee"] = data.TotalFee
	m["out_trade_no"] = data.OutTradeNo
	m["nonce_str"] = data.NonceStr
	data.Sign = wxpayCalcSign(m, wxPayApiKey) //这个是计算wxpay签名的函数上面已贴出

	bytesReq, err := xml.Marshal(data)
	if err != nil {
		return nil, logs.SysErr(err)
	}

	strReq := string(bytesReq)
	strReq = strings.Replace(strReq, "UnifyOrderReq", "xml", -1)
	bytesReq = []byte(strReq)

	//发送unified order请求.统一下单接口
	url := "https://api.mch.weixin.qq.com/pay/unifiedorder"
	req, err := http.NewRequest("POST", url, bytes.NewReader(bytesReq))
	if err != nil {
		return nil, logs.SysErr(err)
	}
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Content-Status", "application/xml;charset=utf-8")

	c := http.Client{}
	resp, err := c.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, logs.SysErr(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, logs.SysErr(err)
	}
	//解析响应结果数据
	respData := new(ReturnData)
	err = xml.Unmarshal(body, respData)
	if err != nil {
		return nil, logs.SysErr(err)
	}
	code := strings.ToUpper(respData.ReturnCode)
	if code == SUCCESS {
		//返回给客户端数据，用于调起支付
		resPayData := new(ResPayData)
		resPayData.TimeStamp = strconv.FormatInt(time.Now().Unix(), 10) //时间戳
		resPayData.PrepayId = respData.PrepayId                         //统一下单接口返回的 prepay_id 参数值
		resPayData.NonceStr = respData.NonceStr                         //随机数
		//签名算法
		resMap := make(map[string]interface{}, 0)
		resMap["appId"] = appid
		resMap["nonceStr"] = respData.NonceStr
		resMap["package"] = "prepay_id=" + respData.PrepayId
		resMap["signType"] = "MD5"
		resMap["timeStamp"] = strconv.FormatInt(time.Now().Unix(), 10)
		resPayData.PaySign = wxpayCalcSign(resMap, wxPayApiKey) //签名
		return resPayData, nil
	} else {
		return nil, logs.SysErr(fmt.Errorf(respData.ReturnMsg))
	}
}

//wxpay计算签名的函数
func wxpayCalcSign(mReq map[string]interface{}, key string) (sign string) {
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

	//STEP3, 在键值对的最后加上key=API_KEY
	if key != "" {
		signStrings = signStrings + "key=" + key
	}

	//STEP4, 进行MD5签名并且将所有字符转为大写.
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	return upperSign
}

//随机字符串
func nonceStr() string {
	return fmt.Sprintf("%s%d", time.Now().Format("20060102150405"), util.RandInt(0000, 9999))
}
