package checkstore

import (
	"bluekinghealth/internal/store"
	"context"
	"github.com/sirupsen/logrus"
)

//实现check子接口

/*
//存储层子接口,这里是check子接口
type CheckCmd interface {
	PaasCheckCmd(ctx context.Context)error
	CmdbCheckCmd(ctx context.Context)error
}
*/
type CheckComponent struct {
	name string
}

func (p *CheckComponent) PaasCheckCmd(ctx context.Context) ([]string, error) {
	var data = []string{"check data..."}

	logrus.Info("finally,use check store  implement func to exec and return data")
	return data, nil
}

func (p *CheckComponent) CmdbCheckCmd(ctx context.Context) ([]string, error) {
	logrus.Info("check impl CmdbCheckCmd")
	return nil, nil
}

func newCheckComponent(name string) *CheckComponent {
	logrus.Info("build new check store func  success and return... ")
	return &CheckComponent{name: name}
}

var _ store.CheckCmd = &CheckComponent{}
