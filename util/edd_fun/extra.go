package edd_fun

import (
	"github.com/gin-gonic/gin"
	"crypto/md5"
	"io"
	"fmt"
	"time"
	"net/http"
)

func GetIp(c *gin.Context) string {
	//ip :=c.Request.RemoteAddr
	//return fmt.Sprintf(ip[0:strings.LastIndex(ip,":")])

	//X-Real_IP是根据nginx的配置的header来的，用于获取客户端的真实信息
	return c.ClientIP()
}

func Md5(value string) string {
	h := md5.New()
	io.WriteString(h, value)
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}

func SetCookie(c *gin.Context, key, value string, expiresTime time.Duration) {
	expires := time.Now().Add(expiresTime)

	cookie := &http.Cookie{
		Name:     key,
		Value:    value,
		Path:     "/",
		HttpOnly: false,
		Expires:  expires,
	}
	http.SetCookie(c.Writer, cookie)
}

func GetCookie(c *gin.Context, key string) (string, error) {
	v, err := c.Request.Cookie(key)
	if err != nil {
		return "", err
	}
	return v.Value,nil
}
