package housekeeper

import(
	"github.com/op/go-logging"
)

type S_configuration struct {
	MQTT struct {
		Broker string `json:"broker"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"mqtt"`
	Webserver struct {
		Protocol string `json:"protocol"`
		Listen string `json:"listen"`
		WebPath string `json:"web_path"`
		Certificate string `json:"certificate"`
		CertificateKey string `json:"certificate_key"`
//		LogRequests bool `json:"log_requests"`
	} `json:"webserver"`
	LogFile string `json:"log_file"`
}

type sharedInformation struct {
	Logger *logging.Logger
	Configuration *S_configuration
	Hub *S_Hub
}

var SharedInformation = sharedInformation{ nil, nil, nil }
