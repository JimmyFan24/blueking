package check

import (
	srvv1 "bluekinghealth/internal/service/v1"
	"bluekinghealth/internal/store"
	"github.com/sirupsen/logrus"
)

type CheckController struct {
	srv srvv1.Service
}
func NewCheckController(component store.CmdFactory)*CheckController{
	logrus.Info("building new check controller with check store factory")
<<<<<<< HEAD
	logrus.Info("building new check controller with check store factory")
=======
>>>>>>> 20bf7d7c8086156ce39db0c253f52e779279e7b7
	return &CheckController{
		srv:srvv1.NewService(component),
	}
}