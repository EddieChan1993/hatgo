package routers

import (
	"github.com/gin-gonic/gin"
	"hatgo/pkg/setting"
	"hatgo/app/middle"
	"hatgo/app/routers/api/v1"
)

func InitRouter() *gin.Engine {
	gin.SetMode(setting.RunMode)
	r:=gin.New()
	r.Use(gin.Recovery())
	if setting.RunMode == gin.DebugMode {
		r.Use(gin.Logger())
	}
	r.Use(middle.Core, middle.TouchBody)

	r.POST("/login",v1.LoginR)

	apiv1 := r.Group("api/v1")
	{
		apiv1.GET("/get-test", v1.GetTestR)
		apiv1.POST("/add-test", v1.AddTest)
		apiv1.PUT("/test/:id", v1.EditTest)
		apiv1.DELETE("/test/:id", v1.DelTest)
	}
	return r
}
