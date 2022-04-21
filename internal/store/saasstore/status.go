package saasstore

import (
	"bluekinghealth/internal/store"
	cmd "bluekinghealth/pkg/command"
	metav1 "bluekinghealth/pkg/meta/v1"
	"context"
	"github.com/sirupsen/logrus"
	"strings"
)

type SaasStatus struct {
	name string
}

//store impl
func (s *SaasStatus) SaasStatus(ctx context.Context, sg *metav1.SaasGroup) (string, error) {
	//

	basciCommand := " docker ps -a| grep "
	logrus.Info("store saas os init")
	logrus.Infof("saasgroup is :%v", sg.Sg)
	for _, saas := range sg.Sg {
		dockercommands := "/data/install/pcmd.sh -H 10.10.26.70 " + " \"" + basciCommand + saas.SaasName + "\""
		saasStatusResult, err := cmd.ExecCmd(dockercommands)
		if err != nil {
			return "", err
		}
		if strings.Contains(saasStatusResult, "Up") {
			logrus.Info("set saas os to running...")
			saas.SaasRunningInfo = "saas is running"
			saas.SaasStatus = true
		} else {
			saas.SaasRunningInfo = "saas is not running"
			saas.SaasStatus = false
		}

	}

	return "saasStatusResult", nil
}

func newSaasStatus(name string) *SaasStatus {
	return &SaasStatus{
		name: name,
	}
}

var _ store.SaasCmd = &SaasStatus{}
