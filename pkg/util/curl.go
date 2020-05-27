package util

import (
	"bytes"
	"fmt"
	"hatgo/pkg/logs"
	"io/ioutil"
	"net/http"
)

const (
	XMLHeader  = "xml"
	JSONHeader = "json"
)

const (
	POST = "POST"
	GET  = "GET"
)

//请求对象
type ReqParams struct {
	Url    string //地址
	Method string //请求方法
	Header string //请求头 JSON或者XML
	Params []byte //请求参数
}

type reqObj struct {
	req *http.Request
}

//初始请求参数
func (p *ReqParams) InitRequest() (req *reqObj, err error) {
	var reqParams *bytes.Reader
	if p.Params != nil {
		reqParams = bytes.NewReader(p.Params)
	}
	obj := new(reqObj)
	obj.req, err = http.NewRequest(p.Method, p.Url, reqParams)
	if err != nil {
		return nil, logs.SysErr(err)
	}
	if p.Method == POST {
		switch p.Header {
		case JSONHeader:
			obj.req.Header.Set("Content-Type", "application/json;charset=UTF-8")
			break
		case XMLHeader:
			obj.req.Header.Set("Accept", "application/xml")
			obj.req.Header.Set("Content-Status", "application/xml;charset=utf-8")
			break
		default:
			obj.req.Header.Set("Content-Type", "application/json;charset=UTF-8")
		}
	}

	return obj, nil
}

//设置header头
func (obj *reqObj) SetHeader(key, val string) {
	obj.req.Header.Set(key, val)
}

//执行请求
func (obj *reqObj) Do() ([]byte, error) {
	defer func() {
		if er := recover(); er != nil {
			logs.SysErr(fmt.Errorf("%v", er))
		}
	}()
	c := http.Client{}
	resp, err := c.Do(obj.req)
	defer resp.Body.Close()
	if err != nil {
		return nil, logs.SysErr(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, logs.SysErr(err)
	}
	return body, nil
}
