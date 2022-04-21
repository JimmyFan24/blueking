package v1

import "github.com/sirupsen/logrus"

type SaasGroup struct {
	Sg []*Saas
}

type Saas struct {
	SaasName        string
	SaasStatus      bool
	SaasRunningInfo string
	SaasBasicInfo   map[string]string
}

var saasList = []string{
	"bk_monitorv3", "bk_itsm", "bk_sops", "bk_log_search", "bk_iam", "bk_user_manage", "bk_nodeman",
}

func NewSaasGroup() *SaasGroup {
	saasgroup := make([]*Saas, len(saasList))
	for i, saas := range saasList {
		saasgroup[i] = NewSaas(saas)
		logrus.Info(saasgroup)
	}
	return &SaasGroup{Sg: saasgroup}
}
func NewSaas(name string) *Saas {
	return &Saas{
		name,
		false,
		"",
		make(map[string]string, 2),
	}
}
