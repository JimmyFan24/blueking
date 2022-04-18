package checkstore

import (
	"bluekinghealth/internal/store"
	"bluekinghealth/pkg/platform"
	"context"
	"github.com/sirupsen/logrus"
	//cmd "bluekinghealth/pkg/command"
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

func (p *CheckComponent) PaasCheckCmd(ctx context.Context) ([]byte, error) {
	//1.执行脚本
	//paascheckresult,err := cmd.ExecCmd("paas")
	/*if err != nil{
		logrus.Errorf("paascheckcmd exec failed..",err)
		return nil, err
	}*/

	var data = []byte(platform.PaasDataTest)

	logrus.Info("finally,use check store  implement func to exec and return data")
	return data, nil
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
