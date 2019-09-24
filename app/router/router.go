package router

import (
	"github.com/gin-gonic/gin"
	"hatgo/app/middle"
	"hatgo/app/router/test"
	"hatgo/pkg/s"
)

func InitRouter() *gin.Engine {
	gin.SetMode(s.RunMode)
	r := gin.New()
	r.Use(gin.Recovery())
	if s.RunMode == gin.DebugMode {
		r.Use(gin.Logger())
	}
	r.Use(middle.Core, middle.TouchBody)

	r.POST("/login", test.RLogin)

	api := r.Group("/")
	{
		api.POST("get-test", test.RGetTest)
		api.POST("add-test", test.RAddTest)
		//api.POST("upload", api.RUpload)
		api.POST("get-xml", test.GetXml)
	}
	return r
}
