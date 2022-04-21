package os

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (o *OsHealthController) OsBasic(c *gin.Context) {
	data, err := o.srv.OsHealth().OsBasic(c)
	if err != nil {
		return
	}
	logrus.Infof("print disk info in controller:%v:%v", data.OsIfG[0].OsBasic[0].ItemName, data.OsIfG[0].OsBasic[0].ItemStatus)
	c.JSON(200, gin.H{
		"osstatus": data.OsIfG,
	})
}
