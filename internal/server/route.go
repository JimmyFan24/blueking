package server

import (
	"bluekinghealth/internal/controller/v1/health"
	oscon "bluekinghealth/internal/controller/v1/os"
	"bluekinghealth/internal/controller/v1/saas"
	_ "bluekinghealth/internal/store"
	healthcmd "bluekinghealth/internal/store/healthstore"
	oscmd "bluekinghealth/internal/store/oshealth"
	saascmd "bluekinghealth/internal/store/saasstore"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func installController(g *gin.Engine) *gin.Engine {
	//返回check factory
	g.Use(Cors())
	//这里返回的是一个实现了storefactory的实例,而且带命名
	healthStoreIns, _ := healthcmd.GetCheckFactory("health")
	saascomponent, _ := saascmd.GetSaasFactory("saas")
	osHealthIns, _ := oscmd.GetOsHealthFactory("oshealth")
	g.GET("/index", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"123": "123",
		})
	})
	v1 := g.Group("/v1")
	{
		checkroute := v1.Group("/health")
		{
			logrus.Info("route regist /v1/health/paas")
			healthController := health.NewHealthController(healthStoreIns)
			checkroute.GET("/paas", healthController.PaasCheck)

		}
		saasroute := v1.Group("/saas")
		{
			logrus.Info("route regist /v1/saas/os")
			saasController := saas.NewSaasController(saascomponent)
			saasroute.GET("/status", saasController.SaasStatus)
		}
		osroute := v1.Group("/os")
		{
			logrus.Info("route regist /v1/os/basic")
			osController := oscon.NewOsHealthController(osHealthIns)
			osroute.GET("/basic", osController.OsBasic)
		}

	}

	return g
}
