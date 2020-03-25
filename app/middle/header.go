package middle

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"hatgo/pkg/logs"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Core(c *gin.Context) {
	method := c.Request.Method

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", fmt.Sprintf("Content-Type,AccessToken,X-CSRF-Token, Authorization, %s", HTTP_TOKEN))
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")

	//放行所有OPTIONS方法
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	// 处理请求
	c.Next()
}

//记录请求体
func TouchBody(c *gin.Context) {
	requestInfo := fmt.Sprintf("\n")

	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery
	clientIP := c.ClientIP()
	method := c.Request.Method
	statusCode := c.Writer.Status()

	if raw != "" {
		path = path + "?" + raw
	}
	requestInfo += fmt.Sprintf("%3d | %13v |%-7s %s\n", statusCode, clientIP, method, path)
	if c.Request.Method == "POST" {
		var headerInfo string
		//get body
		b, _ := ioutil.ReadAll(c.Request.Body)
		s, _ := url.PathUnescape(string(b))
		headerInfo += fmt.Sprintf("Content-Type:%s\n", c.GetHeader("Content-Type"))
		headerInfo += fmt.Sprintf("%s:%s\n", HTTP_TOKEN, c.GetHeader(HTTP_TOKEN))
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(b))
		requestInfo += fmt.Sprintf("%s%s\n", headerInfo, s)
	}
	requestInfo += fmt.Sprintf("-----------------------------------------------------------------------------")
	logs.LogsReq.Info(requestInfo)
	c.Next()
}
