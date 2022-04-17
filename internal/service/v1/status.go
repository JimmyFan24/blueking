package v1

type StatusSrv interface {
	PaasStatus() string
	CmdbStatus() string
}

type statusSrv struct {
}

func (s statusSrv) PaasStatus() string {
	panic("implement me")
}

func (s statusSrv) CmdbStatus() string {
	panic("implement me")
}

var _ StatusSrv = &statusSrv{}
