package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hatgo/pkg/e"
)

func Auth(c *gin.Context) {
	if cookie,err:=c.Request.Cookie("app-token");err== nil {
		value:=cookie.Value
		if value=="value-01" {
			c.Next()
			return
		}
	}
	authCode:=http.StatusUnauthorized
	c.JSON(authCode,e.ResOutput(authCode,http.StatusText(authCode)))
	c.Abort()
}