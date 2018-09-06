package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"hatgo/pkg/util"
)

func LoginR(c *gin.Context) {
	expires := time.Hour
	util.SetCookie(c, "app-token", "value-01", expires)
	c.JSON(http.StatusOK, "LoginR Ok")
}
