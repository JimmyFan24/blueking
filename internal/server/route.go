package server

import (
	"bluekinghealth/internal/controller/v1/check"
	_ "bluekinghealth/internal/store"
	cmd "bluekinghealth/internal/store/checkstore"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func installController(g *gin.Engine) *gin.Engine {
	//返回check factory
	component, _ := cmd.GetCheckFactory("check")
	g.GET("/index", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"123": "123",
		})
	})
	v1 := g.Group("/v1")
	{
		checkroute := v1.Group("/check")
		{
			logrus.Info("route regist /v1/check/paas")
			checkController := check.NewCheckController(component)
			checkroute.GET("/paas", checkController.PaasCheck)
		}

	}

	return g
}
