package v1

import (
	"bluekinghealth/internal/store"
	"github.com/sirupsen/logrus"
)

type Service interface {
	//Component and service health
	Health() HealthSrv
	//Saas health
	Saas() SaasSrv
	//Os health
	OsHealth() OsSrv
}

func NewService(cmdFactory store.StoreFactory) Service {
	logrus.Info("building new service interface with health factory..")
	return &service{
		cmdFactory,
	}
}

//一个Serice的实现
type service struct {
	CmdFactory store.StoreFactory
}

func (s *service) OsHealth() OsSrv {
	return newOsSrv(s)
}

func (s *service) Saas() SaasSrv {
	return newSaasSrv(s)
}

func (s *service) Health() HealthSrv {
	return newHealthService(s)
}

var _ Service = &service{}
