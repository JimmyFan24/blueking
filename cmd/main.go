package main

import "bluekinghealth/internal/server"

func main() {
	//1.health linux system environment
	//2.health all components os
	/*statusReslut,err:= platform.PlatformHeatlthStatus()
	if err!=nil{

		os.Exit(1)
	}

	_,_ = json.Marshal(statusReslut)

	r := gin.Default()
	r.LoadHTMLGlob("C:\\Users\\jimmy\\GolandProjects\\bluekinghealthz\\pkg\\template\\*")
	r.GET("/index", func(context *gin.Context) {
		context.HTML(200,"index.html",gin.H{
			"Title":"this is blueking health health index page",
		})
	})
	r.GET("/os", func(context *gin.Context) {
		context.HTML(200,"os.html",gin.H{
			"id0":0,
			"id1":1,
			"data0":statusReslut[0],
			"data1":statusReslut[1],
		})
	})
	r.Run(":9999")

	*/
	//3.health saas
	server.NewServer().ServerRun()
	//str := ` 16:24:08 [SUCCESS] 10.10.26.73 paas-apigw(http://10.10.26.73:8005/api/healthz/): true paas-appengine(http://10.10.26.73:8000/v1/healthz/): true paas-esb(http://10.10.26.73:8002/healthz/)   : true paas-login(http://10.10.26.73:8003/healthz/) : true paas-paas(http://10.10.26.73:8001/healthz/)  : true `
	//reg,_ := regexp.Compile(`\n`)
	//str_list := reg.Split(str,-1)
	//fmt.Println(str_list[1])

}
