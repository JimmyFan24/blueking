package check

import (
	"bluekinghealth/pkg/script"
	"fmt"
	"github.com/sirupsen/logrus"
	"sync"

	"os/exec"
)

type OriCheckData struct {
	Check map[string]string
}

func Check()  *OriCheckData {
	scriptdata ,err := healthCheck()
	if err != nil {
		logrus.Errorf("Check failed:%v",err)
		return &OriCheckData{}
	}

	return &OriCheckData{
		scriptdata,
	}
}
//check moudle health with script
func healthCheck()(map[string]string, error){
	moudlelist := []string{
		"bkssm","bkiam", "usermgr", "paas", "cmdb", "gse", "job", "consul", "bkmonitorv3",
	}

	checkresult ,err := moudleHeatlthCheck(moudlelist)


	if err != nil {
		logrus.Errorf("moudlehealthscript check failed:%v",err)
		return nil,err
	}
	return checkresult,nil

}
func ExeCmd(moudle string,wg *sync.WaitGroup,out *map[string]string){
	//fmt.Println(moudle,script.Check_test)
	moudlecheckoutput,err:=exec.Command("/bin/bash","-c",script.Check_test+moudle).Output()
	if err != nil {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
	}

	(*out)[moudle] = string(moudlecheckoutput)
	//fmt.Printf("%s", moudlecheckoutput)

	wg.Done()
}
func moudleHeatlthCheck(moudlelist []string)(map[string]string, error){
	//1.根据传进来的moudle,执行check脚本,获取打印结果
	outputMap := map[string]string{}
	wg := new(sync.WaitGroup)
	for _,m := range moudlelist{
		wg.Add(1)
		go ExeCmd(m,wg,&outputMap)
	}
	wg.Wait()
	fmt.Println(outputMap)
	//moudlecheckoutput,e:=exec.Command("/bin/bash","-c",script.CheckScript,moudle).Output()
	return outputMap,nil


}
