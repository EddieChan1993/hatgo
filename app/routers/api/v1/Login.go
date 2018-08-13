package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	expires:=time.Now().Add(time.Hour)

	cookie:=&http.Cookie{
		Name:"app-token",
		Value:"value-01",
		Path:"/",
		HttpOnly:false,
		Expires:expires,
	}
	http.SetCookie(c.Writer, cookie)
	c.JSON(http.StatusOK,"Login Ok")
}