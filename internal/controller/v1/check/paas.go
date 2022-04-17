package check

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (c *CheckController) PaasCheck(context *gin.Context) {
	//srv 的check返回一个实现了Check接口方法的service实例
	logrus.Info("catch request and controller begin to use sub check service :paas check service")
	data, err := c.srv.Check().PaasCheck(context)
	if err != nil {
		context.JSON(400, "paascheck response json failed")
	}

	context.JSON(200, data[0])
	return
}
