package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hatgo/pkg/e"
)

var Channels = make(map[string]int)

func Auth(c *gin.Context) {
	login := c.PostForm("login")
	pass := c.PostForm("pass")

	if login == "" && pass == "" {
		authCode := http.StatusUnauthorized
		e.Output(c, authCode, http.StatusText(authCode))
		c.Abort()
	}

	if cookie, err := c.Request.Cookie("app-token"); err == nil {
		value := cookie.Value
		if value == "value-01" {
			c.Next()
			return
		}
	}
}
