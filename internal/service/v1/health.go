package v1

import (
	"bluekinghealth/internal/store"
	metav1 "bluekinghealth/pkg/meta/v1"
	"context"
	"github.com/sirupsen/logrus"
)

type HealthSrv interface {
	PaasHealth(ctx context.Context) ([]*metav1.App, error)
}
type healthSrv struct {
	checkCmdFactory store.StoreFactory
}

var _ HealthSrv = &healthSrv{}

func newHealthService(srv *service) *healthSrv {
	logrus.Info("building health service implement whit given health factory")
	return &healthSrv{checkCmdFactory: srv.CmdFactory}
}

//service impl
func (c *healthSrv) PaasHealth(ctx context.Context) ([]*metav1.App, error) {
	//调用store层
	logrus.Info("use paas health service  implement func to use health store ")
	//新建paasapp
	paas := metav1.NewApp("paas")
	_, err := c.checkCmdFactory.Health().PaasCheckCmd(ctx, paas)
	if err != nil {
		logrus.Error("service paascheck failed")
		return nil, err
	}
	logrus.Infof("print com health in service impl :%v,%v", paas[0].Components[0].ComCheck, paas[0].ServerIp)
	//返回的paas实例已经修改状态,可以生成json返回
	var result = map[string]interface{}{}
	//var resultJson []byte
	/*for _, p := range paas {
		resultJson, err = json.Marshal(p)
		if err != nil {
			logrus.Errorf("josn marshal failed:%v", err)
		}
		result = append(result, string(resultJson))
	}*/

	logrus.Infof("service frame :print result json:%s", result)
	return paas, nil
}
