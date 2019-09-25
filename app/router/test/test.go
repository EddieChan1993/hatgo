package test

import (
	"github.com/gin-gonic/gin"
	"hatgo/app/service"
	"hatgo/pkg/util"
)

func RGetTest(c *gin.Context) {
	err := service.SGetTestT(c)
	util.Warning(c, err)
}

func RAddTest(c *gin.Context) {
	err := service.SAddTest(c)
	if err != nil {
		util.Warning(c, err)
	} else {
		util.Success(c, "ok")
	}
}

//func RUpload(c *gin.Context) {
//	path, err := service.SUpload(c)
//	if err != nil {
//		e.Warning(c, err)
//	} else {
//		e.Success(c, path)
//	}
//}

func GetXml(c *gin.Context) {
	service.GetXml(c)
}
