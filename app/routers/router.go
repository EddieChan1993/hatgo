package routers

import (
	"github.com/gin-gonic/gin"
	"hatgo/pkg/setting"
	"hatgo/app/routers/api/v1"
	"hatgo/app/middle"
)

func InitRouter() *gin.Engine {
	gin.SetMode(setting.RunMode)
	r := gin.New()
	r.Use(gin.Recovery())
	if setting.RunMode == gin.DebugMode {
		r.Use(gin.Logger())
	}
	r.Use(middle.Core, middle.TouchBody)

	r.POST("/login", v1.RLogin)

	apiv1 := r.Group("api")
	{
		apiv1.POST("/get-test", v1.RGetTest)
		apiv1.POST("/add-test", v1.RAddTest)
		apiv1.POST("/upload", v1.RUpload)
		apiv1.DELETE("/test/:id", v1.RDelTest)
	}
	return r
}
