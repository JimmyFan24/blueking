package v1

import (
	"bluekinghealth/internal/store"
	"github.com/sirupsen/logrus"
)

type Service interface {
	Check() CheckSrv
	Status() StatusSrv
}

func NewService(component store.CmdFactory) Service {
	logrus.Info("building new service interface  with check factory..")
	return &service{
		component,
	}
}

type service struct {
	Component store.CmdFactory
}

func (s *service) Check() CheckSrv {
	return newCheckService(s)
}

func (s service) Status() StatusSrv {
	panic("implement me")
}

var _ Service = &service{}

func newService() Service {
	return &service{}
}
