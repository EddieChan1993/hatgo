package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"hatgo/pkg/util/edd_fun"
)

func Login(c *gin.Context) {
	expires:=time.Hour
	edd_fun.SetCookie(c,"app-token","value-01",expires)
	c.JSON(http.StatusOK,"Login Ok")
}