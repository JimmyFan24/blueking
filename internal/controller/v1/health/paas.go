package health

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (c *HealthController) PaasCheck(context *gin.Context) {
	//srv 的check返回一个实现了Check接口方法的service实例
	logrus.Info("catch request and controller begin to use sub health service :paas health service")
	data, err := c.srv.Health().PaasHealth(context)
	if err != nil {
		context.JSON(400, "paascheck response json failed")
	}
	//context.String(200,data)
	logrus.Infof("controller frame print com health:%v", data[0].Components[0].ComCheck)
	context.JSON(200, gin.H{
		"result": data,
	})
	return
}
