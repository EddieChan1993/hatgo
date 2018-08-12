package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hatgo/pkg/e"
	"hatgo/service"
)

func GetTest(c *gin.Context) {
	err := service.GetTestT(c)
	c.JSON(http.StatusOK, e.ResWarning(err.Error()))
}

func AddTest(c *gin.Context) {

}

func EditTest(c *gin.Context) {

}

func DelTest(c *gin.Context) {

}
