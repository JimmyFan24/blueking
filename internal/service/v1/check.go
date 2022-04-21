package v1

import (
	"bluekinghealth/internal/store"
	metav1 "bluekinghealth/pkg/meta/v1"
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"regexp"
)

type CheckSrv interface {
	PaasCheck(ctx context.Context) ([]string, error)
	CmdbCheck(ctx context.Context) ([]string, error)
}
type checkSrv struct {
	component store.CmdFactory
}

var _ CheckSrv = &checkSrv{}

func newCheckService(srv *service) *checkSrv {
	logrus.Info("building check service implement whit given check factory")
	return &checkSrv{component: srv.Component}
}

//service impl
func (c *checkSrv) PaasCheck(ctx context.Context) ([]string, error) {
	//调用store层
	logrus.Info("use paas check service  implement func to use check store ")
	//新建paasapp
	paas := metav1.NewApp("paas")
	_, err := c.component.Check().PaasCheckCmd(ctx, paas)
	if err != nil {
		logrus.Error("service paascheck failed")
		return nil, err
	}

	//返回的paas实例已经修改状态,可以生成json返回
	var result = []string{}
	var resultJson []byte
	for _, p := range paas {
		resultJson, err = json.Marshal(p)
		if err != nil {
			logrus.Errorf("josn marshal failed:%v", err)
		}
		result = append(result, string(resultJson))
	}

	logrus.Infof("service frame :print result json:%s", string(resultJson))
	return result, nil
}

func (c *checkSrv) CmdbCheck(ctx context.Context) ([]string, error) {
	data, err := c.component.Check().CmdbCheckCmd(ctx)
	if err != nil {
		logrus.Error("service cmdbcheck failed")
		return nil, err
	}
	//拿到脚本执行结果,开始处理数据

	return []string{string(data)}, nil
}
func dataReg(str string) []string {
	reg1, _ := regexp.Compile(`\n`)
	return reg1.Split(str, -1)
}
