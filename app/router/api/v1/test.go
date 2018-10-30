package v1

import (
	"github.com/gin-gonic/gin"
	"hatgo/pkg/e"
	"hatgo/app/service"
)

func RGetTest(c *gin.Context) {
	err := service.SGetTestT(c)
	e.Warning(c, err)
}

func RAddTest(c *gin.Context) {
	err := service.SAddTest(c)
	if err != nil {
		e.Warning(c, err)
	} else {
		e.Success(c, "ok")
	}
}

func RUpload(c *gin.Context) {
	path, err := service.SUpload(c)
	if err != nil {
		e.Warning(c, err)
	} else {
		e.Success(c, path)
	}
}

func GetXml(c *gin.Context) {
	service.GetXml(c)
}
