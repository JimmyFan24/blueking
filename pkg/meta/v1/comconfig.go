package v1

type AppComponent struct {
	PaasComponents []string
}

func NewComponents(appName string) []*Component {
	switch appName {
	case "paas":
		return []*Component{&Component{
			ComName:   "paas-paas",
			ComStatus: false,
			ComCheck:  false,
			ComPort:   "8001",
		},
			&Component{
				ComName:   "paas-login",
				ComStatus: false,
				ComCheck:  false,
				ComPort:   "8003",
			},
			&Component{
				ComName:   "paas-esb",
				ComStatus: false,
				ComCheck:  false,
				ComPort:   "8002",
			},
			&Component{
				ComName:   "paas-apigw",
				ComStatus: false,
				ComCheck:  false,
				ComPort:   "8005",
			}, &Component{
				ComName:   "paas-appengine",
				ComStatus: false,
				ComCheck:  false,
				ComPort:   "8000",
			}}
	default:

	}
	return nil
}
