package command

import (
	metav1 "bluekinghealth/pkg/meta/v1"
	"github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

func BlueKingCmd(app *metav1.App, action string) error {
	//1.获取LanIp
	//lanIp ,_ := ExecCmd("source /data/install/utils.fc && echo  -n $LAN_IP")
	checkAction(app)
	//logrus.Info("---------"+app.ServerIp)
	//logrus.Info("lan ip is "+lanIp)

	return nil
}

func ExecCmd(commands string) (string, error) {
	//commands := Check_test
	//logrus.Info(commands)
	//logrus.Info(appUrl)
	//output ,err := exec.Command("curl","-s","-w","%{http_code}",appUrl).Output()
	output, err := exec.Command("/bin/bash", "-c", commands).Output()
	if err != nil {
		logrus.Errorf("check script  exec failed..")
		return "", err
	}
	logrus.Infof("cmd exec opuput for paas-esb is %v ---->" + string(output) + "<-----")
	return string(output), err
}

func RemoteCmd(commands string, serverIp string) (string, error) {
	//使用蓝鲸本身封装的pcmd.sh

	remoteCommand := "/data/install/pcmd.sh -H " + serverIp + " \"" + commands + "\"" + " |grep -vE \"SUCCESS|FAILURE\""
	logrus.Info("remote command is " + remoteCommand)
	remoteOutput, err := ExecCmd(remoteCommand)
	if err != nil {
		logrus.Errorf("exec remote cmd with pcmd failed:%v", err)
		return "", err
	}

	return remoteOutput, nil
}

//通过传进来的app实例和moudle,拼接命令
func checkAction(app *metav1.App) error {
	//var urlList = make(map[string]string)
	lanIp, _ := ExecCmd("source /data/install/utils.fc && echo  -n $LAN_IP")

	if app.ServerIp == lanIp {
		//logrus.Info("no remote cmd...")
		for _, com := range app.Components {
			requestUrl := "curl -o /dev/null -s -w %{http_code} " + "http://" + app.ServerIp + ":" + (*com).ComPort + "/healthz/" + "\"" + " |grep -vE \"SUCCESS|FAILURE\" "
			//urlList[(*com).ComName] = paasUrlList[requestUrl]
			//logrus.Info("check action :"+requestUrl)
			requestresult, err := ExecCmd(requestUrl)
			if err != nil {
				logrus.Errorf("paascheckcmd exec failed..", err)
				return err
			}
			//logrus.Infof("打印requestresult:%v",string(requestresult))
			//如果返回的状态码包含200,则把组件的状态设置为ture
			if strings.Contains(string(requestresult), "200") {
				//logrus.Info("-----------设置200------------")
				(*com).ComCheck = true
			}
		}

	} else {
		for _, com := range app.Components {
			requestUrl := "curl -o /dev/null -s -w %{http_code} " + "http://" + app.ServerIp + ":" + (*com).ComPort + "/healthz/"
			//urlList[(*com).ComName] = paasUrlList[requestUrl]

			//logrus.Info("check action :"+requestUrl)
			requestresult, err := RemoteCmd(requestUrl, app.ServerIp)
			if err != nil {
				logrus.Errorf("paascheckcmd exec failed..", err)
				return err
			}
			//logrus.Infof("打印requestresult:%v",string(requestresult))
			//如果返回的状态码包含200,则把组件的状态设置为ture
			if strings.Contains(string(requestresult), "200") {
				logrus.Info("-----------设置200------------")
				(*com).ComCheck = true
			}
		}
	}

	return nil
}
