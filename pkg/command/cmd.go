package command

import (
	"github.com/sirupsen/logrus"
	"os/exec"
)


func ExecCmd(moudle string)([]byte,error)  {
	commands := CheckScript+moudle
	output ,err := exec.Command("/bin/bash","-c",commands).Output()
	if err != nil{
		logrus.Errorf("check script  exec failed..")
		return nil, err
	}
	return output, err
}