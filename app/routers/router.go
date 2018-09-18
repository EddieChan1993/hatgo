package routers

import (
	"github.com/gin-gonic/gin"
	"hatgo/pkg/conf"
	"hatgo/app/routers/api/v1"
	"hatgo/app/middle"
)

func InitRouter() *gin.Engine {
	gin.SetMode(conf.RunMode)
	r := gin.New()
	r.Use(gin.Recovery())
	if conf.RunMode == gin.DebugMode {
		r.Use(gin.Logger())
	}
	r.Use(middle.Core, middle.TouchBody)

	r.POST("/login", v1.RLogin)

	api := r.Group("/")
	{
		api.POST("get-test", v1.RGetTest)
		api.POST("add-test", v1.RAddTest)
		api.POST("upload", v1.RUpload)
		api.DELETE("test/:id", v1.RDelTest)
	}
	return r
}
