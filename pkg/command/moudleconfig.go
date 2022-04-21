package command

import (
	"github.com/sirupsen/logrus"
	"strings"
)

func GetMoudleIpByEnv(name string) (string, error) {
	//1.拼接ip查询命令
	ipCommand := "source /data/install/utils.fc" + "&&" + "echo -n $BK_" + strings.ToUpper(name) + "_IP"
	logrus.Info("get ip cmd is " + ipCommand)
	ipList, err := ExecCmd(ipCommand)
	if err != nil {
		logrus.Errorf("get moudle from bk env failed:%v", err)
		return "", err
	}
	return string(ipList), nil
}
