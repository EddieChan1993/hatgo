package middle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hatgo/pkg/e"
)
func Auth(c *gin.Context) {
	authCode := http.StatusUnauthorized
	token := c.GetHeader("token")
	if token == "" {
		e.Output(c, authCode, http.StatusText(authCode))
		c.Abort()
		return
	}
	c.Set("uid", 12)
	//c.GetInt64("uid")
	c.Next()
}
