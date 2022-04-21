package v1

type PaasApp struct {
	Name       string       `json:"paas"`
	Components []*Component `json:"components"`
	ServerIp   string       `json:"serverip"`
}

/*type Component struct {
	ComName string `json:"comname"`
	ComStatus bool `json:"comstatus"`
	ComCheck bool `json:"comcheck"`
	ComPort string `json:"comport"`
}*/
/*func NewApp() *PaasApp{
	serverIp ,_ := cmd.GetMoudleIpByEnv("iam")
	logrus.Info("server ip is "+serverIp)
	app := &PaasApp{
		Name:       "paas",
		ServerIp:   serverIp,
		Components:[]*Component{
			&Component{
				"usermgr",false,false,"8009",
			},
			&Component{
				"bkssm",false,false,"5000",
			},
			&Component{
				"bkiam",false,false,"5001",
			},
		},
	}

	return app
}*/
