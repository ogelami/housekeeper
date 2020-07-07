package housekeeper

import (
	"github.com/op/go-logging"
)

type configuration struct {
	MQTT struct {
		Broker   string `json:"broker"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"mqtt"`
	Webserver struct {
		Protocol       string `json:"protocol"`
		Listen         string `json:"listen"`
		WebPath        string `json:"web_path"`
		Certificate    string `json:"certificate"`
		CertificateKey string `json:"certificate_key"`
		//		LogRequests bool `json:"log_requests"`
	} `json:"webserver"`
	LogFile string `json:"log_file"`
}

//Logger used by housekeeper to print to console and log file.
var Logger *logging.Logger

//Configuration parsed configuration struct.
var Configuration *configuration

//Hub webserver hub for keeping track of connected clients.
var Hub *S_Hub
