package v1

import (
	"bluekinghealth/internal/store"
	metav1 "bluekinghealth/pkg/meta/v1"
	"context"
)

type OsSrv interface {
	OsBasic(c context.Context) (*metav1.OsInfoGroup, error)
}
type osSrv struct {
	osCmd store.StoreFactory
}

//Os service impl
func (o *osSrv) OsBasic(c context.Context) (*metav1.OsInfoGroup, error) {
	osBasic := metav1.NewOsInfoGroup()
	_, err := o.osCmd.OsHealth().Basic(c, osBasic)
	if err != nil {
		return nil, err
	}

	return osBasic, nil
}

func newOsSrv(srv *service) *osSrv {
	return &osSrv{
		osCmd: srv.CmdFactory,
	}
}

var _ OsSrv = &osSrv{}
