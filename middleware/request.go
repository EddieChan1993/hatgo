package middleware

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"bytes"
	"net/url"
	"fmt"
	"hatgo/logging"
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

		//get header
		h := c.Request.Header
		//get body
		b, _ := ioutil.ReadAll(c.Request.Body)
		s, _ := url.PathUnescape(string(b))

		contentType, has := h["Content-Type"]
		if has {
			headerInfo += fmt.Sprintf("%s\n", contentType[0])
		}
		cookie, has := h["Cookie"]
		if has {
			headerInfo += fmt.Sprintf("%s\n", cookie[0])
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(b))
		requestInfo += fmt.Sprintf("%s%s\n", headerInfo, s)
	}
	requestInfo += fmt.Sprintf("-----------------------------------------------------------------------------")
	logging.Logs.Info(requestInfo)
	c.Next()
}
