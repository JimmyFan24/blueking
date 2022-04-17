package v1

import (
	"bluekinghealth/internal/store"
	"context"
	"github.com/sirupsen/logrus"
)

type CheckSrv interface {
	PaasCheck(ctx context.Context) ([]string,error)
	CmdbCheck(ctx context.Context) ([]string,error)
}
type checkSrv struct {
	component store.CmdFactory
}

var _ CheckSrv = &checkSrv{}
func newCheckService(srv *service)*checkSrv{
	logrus.Info("building check service implement whit given check factory")
	return &checkSrv{component: srv.Component}
}
func (c *checkSrv) PaasCheck(ctx context.Context) ([]string,error ){
	//调用store层
	logrus.Info("use paas check service  implement func to use check store ")
	data,err := c.component.Check().PaasCheckCmd(ctx)
	if err!=nil{
		logrus.Error("service paascheck failed")
		return nil,err
	}

	logrus.Infof("service paascheck success,and the data is :%v",data[0])
	return data,nil
}

func (c *checkSrv) CmdbCheck(ctx context.Context) ([]string,error) {
	return nil,nil
}

