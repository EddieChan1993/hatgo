package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"hatgo/pkg/util"
)

func RLogin(c *gin.Context) {
	expires := time.Hour
	util.SetCookie(c, "app-token", "value-01", expires)
	c.JSON(http.StatusOK, "RLogin Ok")
}
