package check

import (
	srvv1 "bluekinghealth/internal/service/v1"
	"bluekinghealth/internal/store"
	"github.com/sirupsen/logrus"
)

type CheckController struct {
	srv srvv1.Service
}

func NewCheckController(component store.CmdFactory) *CheckController {
	logrus.Info("building new check controller with check store factory")
	return &CheckController{
		srv: srvv1.NewService(component),
	}
}
