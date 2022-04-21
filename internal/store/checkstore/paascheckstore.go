package checkstore

import (
	"bluekinghealth/internal/store"
	cmd "bluekinghealth/pkg/command"
	metav1 "bluekinghealth/pkg/meta/v1"
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

func (p *CheckComponent) PaasCheckCmd(ctx context.Context, app []*metav1.App) (string, error) {
	for _, a := range app {
		logrus.Info("store cmd check loop")
		logrus.Info(a)
		err := cmd.BlueKingCmd(a, "check")
		if err != nil {
			return "", err
		}
		logrus.Info("finally,use check store  implement func to exec and return data")

	}
	return "", nil
}

func (p *CheckComponent) CmdbCheckCmd(ctx context.Context) ([]byte, error) {
	var data = "check cmdb data..."

	//logrus.Info("finally,use check store  implement func to exec and return data")
	return []byte(data), nil
}

func newCheckComponent(name string) *CheckComponent {
	logrus.Info("build new check store func  success and return... ")
	return &CheckComponent{name: name}
}

var _ store.CheckCmd = &CheckComponent{}
