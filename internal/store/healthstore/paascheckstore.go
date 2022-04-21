package healthstore

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
type HealthCmd interface {
	PaasCheckCmd(ctx context.Context)error
	CmdbCheckCmd(ctx context.Context)error
}
*/
type CheckComponent struct {
	name string
}

func (p *CheckComponent) PaasCheckCmd(ctx context.Context, app []*metav1.App) (string, error) {
	for _, a := range app {
		logrus.Info("store cmd health loop")

		err := cmd.BlueKingCmd(a, "health")
		if err != nil {
			return "", err
		}
		logrus.Info("finally,use health store  implement func to exec and return data")
		logrus.Info(a.Components[0].ComCheck)
	}
	logrus.Infof("print on store frame  comcheck value:%v,%v", app[0].ServerIp, app[0].Components[0].ComCheck)
	return "", nil
}

func (p *CheckComponent) CmdbCheckCmd(ctx context.Context) ([]byte, error) {
	var data = "health cmdb data..."

	//logrus.Info("finally,use health store  implement func to exec and return data")
	return []byte(data), nil
}

func newCheckComponent(name string) *CheckComponent {
	logrus.Info("build new health store func  success and return... ")
	return &CheckComponent{name: name}
}

var _ store.HealthCmd = &CheckComponent{}
