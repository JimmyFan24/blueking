package command

import (
	metav1 "bluekinghealth/pkg/meta/v1"
	"github.com/sirupsen/logrus"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

const LANIP = "10.10.26.69"

func BlueKingCmd(app *metav1.App, action string) error {
	//1.获取LanIp
	//lanIp ,_ := ExecCmd("source /data/install/utils.fc && echo  -n $LAN_IP")
	//logrus.Infof("before Bluekingcmd com health:%v,%v",app.Components[0].ComCheck,app.ServerIp)
	checkAction(app)
	statusAction(app)
	//logrus.Infof("after Bluekingcmd com health:%v,%v",app.Components[0].ComCheck,app.ServerIp)
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
		logrus.Errorf("health script  exec failed..")
		return "", err
	}
	//logrus.Infof("cmd exec opuput for paas-esb is %v ---->" + string(output) + "<-----")
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
			requestUrl := "curl -o /dev/null -s -w %{http_code} " + "http://" + app.ServerIp + ":" + (*com).ComPort + (*com).ComCheckPath + "\"" + " |grep -vE \"SUCCESS|FAILURE\" "
			//urlList[(*com).ComName] = paasUrlList[requestUrl]
			//logrus.Info("health action :"+requestUrl)
			requestresult, err := ExecCmd(requestUrl)
			if err != nil {
				logrus.Errorf("paascheckcmd exec failed..", err)
				return err
			}
			//logrus.Infof("打印requestresult:%v",string(requestresult))
			//如果返回的状态码包含200,则把组件的状态设置为ture
			logrus.Infof("print request result :" + string(requestresult))
			if strings.Contains(string(requestresult), "200") {
				//logrus.Info("-----------设置200------------")
				(*com).ComCheck = true
				(*com).ComStatus = true
			} else {
				logrus.Info("设置com health false")
				(*com).ComCheck = false
			}
		}

	} else {
		for _, com := range app.Components {
			requestUrl := "curl -o /dev/null -s -w %{http_code} " + "http://" + app.ServerIp + ":" + (*com).ComPort + (*com).ComCheckPath
			//urlList[(*com).ComName] = paasUrlList[requestUrl]

			//logrus.Info("health action :"+requestUrl)
			requestresult, err := RemoteCmd(requestUrl, app.ServerIp)
			if err != nil {
				logrus.Errorf("paascheckcmd exec failed..", err)
				return err
			}
			logrus.Infof("打印requestresult:%v", string(requestresult))
			//如果返回的状态码包含200,则把组件的状态设置为ture
			if strings.Contains(string(requestresult), "200") {
				logrus.Info("-----------设置200------------")
				(*com).ComCheck = true
				(*com).ComStatus = true
			} else {
				logrus.Info("设置com health false")

				(*com).ComCheck = false
			}

		}

	}
	logrus.Infof("print on cmd frame  comcheck value:%v:%v", app.ServerIp, app.Components[0].ComCheck)
	return nil
}
func statusAction(app *metav1.App) error {
	commands := "/data/install/bkcli os " + app.Name + " " + app.ServerIndex

	for _, com := range app.Components {
		statusComand := commands + "| grep " + (*com).ComName
		//执行状态查询
		logrus.Infof("os cmd is %v", statusComand)
		statusResult, err := ExecCmd(statusComand)
		logrus.Infof("print com os in os action :%v", statusResult)
		if err != nil {
			logrus.Infof("os cmd exec faild:%v", err)
			return err
		}
		if strings.Contains(statusResult, "inactive") {
			logrus.Infof("set com :%v os in false...", (*com).ComName)
			(*com).ComStatus = false
		} else {
			logrus.Infof("set com :%v os in true...", (*com).ComName)
			(*com).ComStatus = true
		}
		statusComand = ""
	}

	return nil
}

func OsHealth(osinfo *metav1.OsInfoGroup) error {

	for _, osItem := range osinfo.OsIfG {
		//logrus.Infof("infocmd is :%v",infocmd)

		if osItem.OsIp == LANIP {
			for _, infoitem := range osItem.OsBasic {

				data, err := ExecCmd(infoitem.ItemCmd)
				if err != nil {
					return err
				}
				//infoitem.ItemResult = data
				if infoitem.ItemName == "diskspace" {
					diskReg := regexp.MustCompile(`\d+%`)
					infoitem.ItemResult = diskReg.FindStringSubmatch(data)[0]
					//logrus.Infof("infoitem.ItemResult is %v",infoitem.ItemResult )
					diskmeg, _ := strconv.Atoi(infoitem.ItemResult)
					//logrus.Info(diskmeg)
					if diskmeg > 90 {
						infoitem.ItemStatus = false
					} else {
						//logrus.Info("set infoitem itemstatus true..")
						infoitem.ItemStatus = true
					}
				} else if strings.Contains(data, infoitem.ItemJud) {
					infoitem.ItemStatus = true
					infoitem.ItemResult = data
				} else {
					infoitem.ItemStatus = false
				}

			}
		} else {
			for _, infoitem := range osItem.OsBasic {
				//remoteCmd :=  "/data/install/pcmd.sh -H " + osItem.OsIp + " \"" + infoitem.ItemCmd + "\""
				data, err := RemoteCmd(infoitem.ItemCmd, osItem.OsIp)
				//RemoteCmd(infoitem.ItemCmd,osItem.OsIp)
				if err != nil {
					return err
				}
				//infoitem.ItemResult = data
				if infoitem.ItemName == "diskspace" {
					diskReg := regexp.MustCompile(`\d+%`)
					infoitem.ItemResult = diskReg.FindStringSubmatch(data)[0]
					//logrus.Infof("infoitem.ItemResult is %v",infoitem.ItemResult )
					diskmeg, _ := strconv.Atoi(infoitem.ItemResult)
					//logrus.Info(diskmeg)
					if diskmeg > 90 {
						infoitem.ItemStatus = false
					} else {
						//logrus.Info("set infoitem itemstatus true..")
						infoitem.ItemStatus = true
					}
				} else if strings.Contains(data, infoitem.ItemJud) {
					infoitem.ItemStatus = true
					infoitem.ItemResult = data
				} else {
					infoitem.ItemStatus = false
				}

			}
		}
	}
	return nil
}
