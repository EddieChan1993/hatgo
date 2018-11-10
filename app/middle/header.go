package middle

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"bytes"
	"net/url"
	"fmt"
	"hatgo/pkg/logs"
)

func Core(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
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
		headerInfo += fmt.Sprintf("%s\n", c.GetHeader("Content-Type"))
		headerInfo += fmt.Sprintf("%s\n", c.GetHeader(HTTP_TOKEN))
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(b))
		requestInfo += fmt.Sprintf("%s%s\n", headerInfo, s)
	}
	requestInfo += fmt.Sprintf("-----------------------------------------------------------------------------")
	logs.LogsReq.Info(requestInfo)
	c.Next()
}
