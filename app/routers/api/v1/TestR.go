package v1

import (
	"github.com/gin-gonic/gin"
	"hatgo/pkg/e"
	"hatgo/app/service"
)

func RGetTest(c *gin.Context) {
	err := service.SGetTestT(c)
	e.Waring(c, err)
}

func RAddTest(c *gin.Context) {

}

func REditTest(c *gin.Context) {

}

func RDelTest(c *gin.Context) {

}
