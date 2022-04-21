package saasstore

import (
	"bluekinghealth/internal/store"
	"github.com/sirupsen/logrus"
)

type saasFactory struct {
	name string
}

func (s saasFactory) OsHealth() store.OsHealthCmd {
	panic("implement me")
}

func (s saasFactory) Health() store.HealthCmd {
	panic("implement me")
}

func (s saasFactory) Saas() store.SaasCmd {
	return newSaasStatus("Saas Status Cmd")
}

var _ store.StoreFactory = &saasFactory{}
var (
	cf store.StoreFactory
)

//返回实现了store.factory实例checkFactory
func GetSaasFactory(name string) (store.StoreFactory, error) {
	logrus.Info("geting new health factory ")
	cf = &saasFactory{
		name: name,
	}
	return cf, nil
}
