package platform

/*func PlatformHeatlthStatus()([]string, error){
	//1.执行status脚本,获取打印结果
	_,e:=exec.Command("/bin/sh","-c",`/data/install/bkcli status all|grep -E "inactive|SUCCESS|deactivating"`).Output()
	if e!=nil{
		logrus.Errorf("exec shell command failed:%v",e)
		return nil, e
		os.Exit(1)
	}

	reg2 := regexp.MustCompile(`\[\d\]`)
	statusResultMap := reg2.Split(status,-1)

	//statusResultMap := reg2.Split(string(o),-1)
	//fmt.Print("--->"+statusResultMap[3]+"<---")
	var finalStatusMap = []string{}
	for i:=0;i<len(statusResultMap);i++{
		//var tmp string
		if match,_ := regexp.MatchString(`\d\n$|^\n$|^$`,statusResultMap[i]);!match{

			finalStatusMap = append(finalStatusMap,statusResultMap[i])
		}

	}


	return finalStatusMap , nil



}*/
/*func PlatformHeatlthCheck() (map[string]string,error) {
	reg2 := regexp.MustCompile(`\[\d\]`)
	statusResultMap := reg2.Split(status,-1)

	//statusResultMap := reg2.Split(string(o),-1)
	//fmt.Print("--->"+statusResultMap[3]+"<---")
	var finalStatusMap = []string{}
	for i:=0;i<len(statusResultMap);i++{
		//var tmp string
		if match,_ := regexp.MatchString(`\d\n$|^\n$|^$`,statusResultMap[i]);!match{

			finalStatusMap = append(finalStatusMap,statusResultMap[i])
		}

	}


	return finalStatusMap , nil

}
*/
/*func openSourceComponent() (string,error){
	//fmt.Println("opensource components hearth checking...")
	return "opensource components hearth checking...",nil
}
func blueKingComponent(str []string)(string,error){
	blueKingComCoreStatus(str)
	//blueKingComCoreHealth()
	//fmt.Println("blueking components hearth checking...")
	return "",nil
}
*/
