package test

import (
	"testing"
	"log"
	"encoding/json"
	"fmt"
)

type Luosimao struct {
	Error int `json:"error"`
	Msg string `json:"msg"`
}

type BaiduOrcToken struct {
	AccessToken string `json:"access_token"`
	SessionKey string `json:"session_key"`
}
//螺丝帽短信接口
func luosimaoApi() {
	url := "http://sms-api.luosimao.com/v1/send.json"
	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": BasicAuth("api", "78aac6166f2318bd2eaceae0fba6aa84"),
	}
	postData := map[string]string{
		"mobile":  "18380591566",
		"message": "go-lang test【环球娃娃】",
	}
	req := NewRequst(url)
	result, err := req.
		SetHeader(headers).
		SetParms(postData).
		Post()
	if err != nil {
		log.Fatalln(err)
	}
	luosimao:=new(Luosimao)
	json.Unmarshal(result,luosimao)
	fmt.Println(luosimao.Msg)
}
//百度搜索
func baiduSearch() {
	searchURL := "http://www.baidu.com/s"
	params:=map[string]string{
		"wd":"星巴克生在美国的那个城市",
	}
	headers:=map[string]string{
		"Content-Type":"application/x-www-form-urlencoded",
	}
	req,err:=NewRequst(searchURL).SetHeader(headers).SetParms(params).Get()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(req))
}
//百度ocr
func baiduOcrToken()  {
	searchURL:= "https://aip.baidubce.com/oauth/2.0/token"
	headers:=map[string]string{
		"Content-Type":"application/json;charset=utf-8",
	}
	params:=map[string]string{
		"grant_type":"client_credentials",
		"client_id":"oIRXqBOdrQm4T5Kd6TAlxEjz",
		"client_secret":"YEynXEaBRqbO0P4AOpywts7r5MLSC2Rb",
	}
	req,err:=NewRequst(searchURL).SetHeader(headers).SetParms(params).Post()
	if err != nil {
		fmt.Println(err)
	}

	baiduToken:=new(BaiduOrcToken)
	json.Unmarshal(req,baiduToken)
	fmt.Println(baiduToken.AccessToken)
}
func TestRequest_Post(t *testing.T) {
	//luosimaoApi()
	baiduSearch()
	//baiduOcrToken()
}