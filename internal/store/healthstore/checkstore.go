package healthstore

import (
	"bluekinghealth/internal/store"
	"github.com/sirupsen/logrus"
)

//实现Store接口
/*
type StoreFactory interface {
	Health() HealthCmd
}

*/
var (
	cf store.StoreFactory
)

type checkFactory struct {
	//store.StoreFactory
	name string
}

func (p *checkFactory) OsHealth() store.OsHealthCmd {
	return nil
}

func (p *checkFactory) Health() store.HealthCmd {
	logrus.Info("building sub health factory :health factory  ")
	return newCheckComponent("new health factory instances...")
}

func (p *checkFactory) Saas() store.SaasCmd {
	//logrus.Info("building sub saas factory :saas factory  ")
	//return newSaasComponent("new saas factory instances...")
	return nil
}

//返回实现了store.factory实例checkFactory
func GetCheckFactory(name string) (store.StoreFactory, error) {
	logrus.Info("geting new health factory ")
	cf = &checkFactory{
		name: name,
	}
	return cf, nil
}
