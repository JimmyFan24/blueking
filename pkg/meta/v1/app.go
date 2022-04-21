package v1

import (
	"github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

type App struct {
	Name       string       `json:"app"`
	Components []*Component `json:"components"`
	ServerIp   string       `json:"serverip"`
}
type Component struct {
	ComName   string `json:"comname"`
	ComStatus bool   `json:"comstatus"`
	ComCheck  bool   `json:"comcheck"`
	ComPort   string `json:"comport"`
}

func NewApp(appName string) []*App {
	//serverIp ,_ := cmd.GetMoudleIpByEnv(appName)
	//要检查是否是只有一个实例
	comands := "source /data/install/utils.fc" + "&&" + "echo -n $BK_" + strings.ToUpper(appName) + "_IP_COMMA"
	//command := "source /data/install/utils.fc"+"&&"+"echo -n $BK_"+strings.ToUpper(appName)+"_IP"
	//处理返回的ip列表
	logrus.Info("find app instances commands is :" + comands)
	serverIpList, err := exec.Command("/bin/bash", "-c", comands).Output()
	if err != nil {
		logrus.Errorf("get moudle server ip faild:%v", err)
	}
	logrus.Info("app server ip : " + string(serverIpList))
	ipList := strings.Split(string(serverIpList), ",")
	com := NewComponents(appName)
	var appList = make([]*App, len(ipList))
	if len(ipList) > 1 {
		logrus.Info("This app is  ha moudle ,which means that it has more than 1 instances")

		for i := 0; i < len(appList); i++ {
			appList[i] = &App{appName, com, ipList[i]}

		}
	} else if len(ipList) == 1 {
		appList[0] = &App{appName, com, ipList[0]}
	}
	logrus.Infof("building new app success...")
	return appList

}
