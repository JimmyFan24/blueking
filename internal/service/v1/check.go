package v1

import (
	"bluekinghealth/internal/store"
	"context"
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
	data, err := c.component.Check().PaasCheckCmd(ctx)
	if err != nil {
		logrus.Error("service paascheck failed")
		return nil, err
	}
	//处理返回的data

	re1 := `\[\d\]`
	re := `[ ]\d{1,}.\d{1,}.\d{1,}.\d{1,}`

	reg,e := regexp.Compile(re1)
	if e != nil{
		logrus.Errorf("re compile failed:%v",err)
	}
	ip_list :=reg.MatchString()

	logrus.Infof("service paascheck success,and the data is :%v", data[0])
	return string(ip_list), nil
}

func (c *checkSrv) CmdbCheck(ctx context.Context) ([]string, error) {
	data,err := c.component.Check().CmdbCheckCmd(ctx)
	if err != nil{
		logrus.Error("service cmdbcheck failed")
		return nil, err
	}
	return data, nil
}
