package util

import (
	"github.com/gin-gonic/gin"
	"crypto/md5"
	"io"
	"fmt"
	"time"
	"net/http"
	"math/rand"
)

//md5加密
func Md5(value string) string {
	h := md5.New()
	io.WriteString(h, value)
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}

//设置cookie
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

//获取cookie
func GetCookie(c *gin.Context, key string) (string, error) {
	v, err := c.Request.Cookie(key)
	if err != nil {
		return "", err
	}
	return v.Value, nil
}

//获取int类型的随机数
func RandInt(start, end int) int {
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	ca := end - start
	return start + rand.Intn(ca)
}

//订单号
func TradeNo(key string) string {
	return fmt.Sprintf("%s%d%d", key, time.Now().UnixNano(), RandInt(1000, 9999)) //订单单号
}
