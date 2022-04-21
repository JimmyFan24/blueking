package command

import (
	"fmt"
	"os/exec"
)

func cmd_test() {
	data, err := exec.Command("/bin/bash", "-c", "ls -l").Output()
	if err != nil {
		return
	}
	fmt.Println(data)
}
