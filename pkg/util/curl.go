package util

import (
	"bytes"
	"hatgo/pkg/logs"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	XMLHeader  = "xml"
	JSONHeader = "json"
)

func GetCurl(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, logs.SysErr(err)
	}
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
	return body, nil
}

/**
	header 默认JSON请求
 */
func PostCurl(url string, params []byte, header string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(params))
	if err != nil {
		return nil, logs.SysErr(err)
	}
	if header == JSONHeader || header == "" {
		req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	} else if header == XMLHeader {
		req.Header.Set("Accept", "application/xml")
		req.Header.Set("Content-Status", "application/xml;charset=utf-8")
	} else {
		return nil, logs.SysErr(fmt.Errorf("未定义请求头"))
	}
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
	return body, nil
}
