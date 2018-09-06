package v1

import (
	"github.com/gin-gonic/gin"
	"hatgo/pkg/e"
	"hatgo/app/service"
)

func GetTestR(c *gin.Context) {
	err := service.GetTestT(c)
	e.Waring(c,err.Error())
}

func AddTest(c *gin.Context) {

}

func EditTest(c *gin.Context) {

}

func DelTest(c *gin.Context) {

}
