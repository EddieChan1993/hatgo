package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hatgo/pkg/e"
	"hatgo/pkg/util/edd_fun"
)

var Channels = make(map[string]int)

func Auth(c *gin.Context) {
	authCode := http.StatusUnauthorized

	login := c.PostForm("login")
	pass := c.PostForm("pass")

	if login == "" && pass == "" {
		e.Output(c, authCode, http.StatusText(authCode))
		c.Abort()
	}

	cookie, err := edd_fun.GetCookie(c, "app_token")
	if err != nil {
		e.Output(c, authCode, http.StatusText(authCode))
		c.Abort()
	} else {
		c.Set("uid", cookie)
		c.Next()
	}

}
