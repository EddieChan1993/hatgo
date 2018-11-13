package v1

import (
	"github.com/gin-gonic/gin"
	"hatgo/app/service"
)

func RWebSocket(c *gin.Context) {
	service.SHandler(c)
}
