package health

import (
	srvv1 "bluekinghealth/internal/service/v1"
	"bluekinghealth/internal/store"
	"github.com/sirupsen/logrus"
)

type HealthController struct {
	srv srvv1.Service
}

func NewHealthController(cmdFactory store.StoreFactory) *HealthController {
	logrus.Info("building new health controller with health store factory")
	return &HealthController{
		srv: srvv1.NewService(cmdFactory),
	}
}
