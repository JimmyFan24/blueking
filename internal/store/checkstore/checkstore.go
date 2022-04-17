package checkstore

import (
	"bluekinghealth/internal/store"
	"github.com/sirupsen/logrus"
)

//实现Store接口
/*
type CmdFactory interface {
	Check() CheckCmd
}

*/
var (
	cf  store.CmdFactory
)

type checkFactory struct {
	//store.CmdFactory
	name string
}

func (p *checkFactory) Check() store.CheckCmd {
	logrus.Info("building sub check factory :check factory  ")
	return newCheckComponent("new check factory instances...")
}

//返回实现了store.factory实例checkFactory
func GetCheckFactory(name string)(store.CmdFactory,error){
	logrus.Info("geting new check factory ")
	cf = &checkFactory{
		name: name,
	}
	return cf,nil
}

