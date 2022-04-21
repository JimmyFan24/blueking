package v1

import (
	"bluekinghealth/internal/store"
	metav1 "bluekinghealth/pkg/meta/v1"
	"context"
	"github.com/sirupsen/logrus"
)

type SaasSrv interface {
	SaasStatus(ctx context.Context) (saas *metav1.SaasGroup, err error)
}
type saasSrv struct {
	saasCmdFactory store.StoreFactory
}

func newSaasSrv(srv *service) *saasSrv {
	return &saasSrv{saasCmdFactory: srv.CmdFactory}
}

//saas service impl
func (s *saasSrv) SaasStatus(ctx context.Context) (saas *metav1.SaasGroup, err error) {
	saasGroup := metav1.NewSaasGroup()
	logrus.Infof("build new saas instances:%v", saasGroup.Sg)
	saasStatus, err := s.saasCmdFactory.Saas().SaasStatus(ctx, saasGroup)

	logrus.Infof("service impl saasstatus :%v", saasStatus)
	return saasGroup, nil

}

var _ SaasSrv = &saasSrv{}
