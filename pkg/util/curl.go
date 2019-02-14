package util

import (
	"bytes"
	"fans/pkg/logs"
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
