package main

import (
	"bluekinghealth/internal/server"
)

func main()  {
	//1.check linux system environment
	//2.check all components status
	/*statusReslut,err:= platform.PlatformHeatlthStatus()
	if err!=nil{

		os.Exit(1)
	}

	_,_ = json.Marshal(statusReslut)

	r := gin.Default()
	r.LoadHTMLGlob("C:\\Users\\jimmy\\GolandProjects\\bluekinghealthz\\pkg\\template\\*")
	r.GET("/index", func(context *gin.Context) {
		context.HTML(200,"index.html",gin.H{
			"Title":"this is blueking health check index page",
		})
	})
	r.GET("/status", func(context *gin.Context) {
		context.HTML(200,"status.html",gin.H{
			"id0":0,
			"id1":1,
			"data0":statusReslut[0],
			"data1":statusReslut[1],
		})
	})
	r.Run(":9999")

*/
	//3.check saas
	server.NewServer().ServerRun()

}
