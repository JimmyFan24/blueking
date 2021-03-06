package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	Addr string
	//Port string
	*gin.Engine
}

func NewServer() *Server {

	return &Server{
		Addr: "127.0.0.1:9999",
	}
}

const cmdb = `
[1] 16:38:17 [FAILURE] 10.10.26.75 Exited with error code 14
cmdb-admin(http://10.10.26.75:9000/healthz)  : false Reason: connection refused
cmdb-api(http://10.10.26.75:9001/healthz)    : false Reason: connection refused
cmdb-auth(http://10.10.26.75:9002/healthz)   : false Reason: connection refused
cmdb-cache(http://10.10.26.75:9014/healthz)  : false Reason: connection refused`

func (s *Server) Run() {

	//_, _ = health.HealthCheck()
	check_test := map[string]string{
		"bkiam":       "bkiam",
		"bkmonitorv3": "bkmonitorv3",
		"paas":        "paas",
	}
	fmt.Println(check_test["paas"])
	r := gin.Default()
	r.Use(Cors())
	r.LoadHTMLGlob("C:\\Users\\jimmy\\GolandProjects\\bluekinghealthz\\pkg\\template\\*")
	r.GET("/index", func(context *gin.Context) {
		context.HTML(200, "index.html", gin.H{
			"Title": "this is blueking health health index page",
		})
	})
	r.GET("/health", func(context *gin.Context) {
		context.HTML(200, "os.html", gin.H{
			"id0":   0,
			"id1":   1,
			"data0": cmdb,
		})
	})
	r.Run(s.Addr)
}
func (s *Server) ServerRun() {
	s.Engine = gin.New()
	installEngine := installController(s.Engine)
	installEngine.Run(":9999")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}

}
