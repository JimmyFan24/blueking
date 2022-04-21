package saas

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (s *SaasController) SaasStatus(c *gin.Context) {
	//调用service
	status, err := s.srv.Saas().SaasStatus(c)
	if err != nil {
		logrus.Errorf("saas os controller failed:%v", err)
	}
	c.JSON(200, gin.H{
		"status": status.Sg,
	})
	logrus.Infof("saas os controller is %v", status.Sg[0].SaasRunningInfo)
	return
}
