package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"./housekeeper"
	"github.com/op/go-logging"
)

/**

make dep
go run -ldflags "-X main.CONFIGURATION_PATH=bin/cfdnsu.conf" main.go
openssl ecparam -genkey -name secp384r1 -out bin/server.key
openssl req -new -x509 -sha256 -key bin/server.key -out bin/server.crt -days 365

* lets keep the CF api calls down by only calling getCFDNSRecordDetails on startup
* upon failure of retrieving the servers ip address, retry in the next cycle
* Shutdown event not firing
* ^C interrupt => runtimer error

* working test check "https://api.ipify.org/"
* ip resolution does not validate the ip addr comming back, if its html or ipv6 it will just be passed to cloudflare which will break.

*/

var (
	CONFIGURATION_PATH string
)

func loadConfiguration() error {
	configurationData, err := ioutil.ReadFile(CONFIGURATION_PATH)

	/*	housekeeper.SharedInformation.Logger.Critical(configurationData)
		housekeeper.SharedInformation.Logger.Critical(CONFIGURATION_PATH)*/

	if err != nil {
		housekeeper.SharedInformation.Logger.Critical(err)
		return err
	}

	err = json.Unmarshal(configurationData, &housekeeper.SharedInformation.Configuration)

	//	housekeeper.SharedInformation.Logger.Critical(configuration)

	if err != nil {
		housekeeper.SharedInformation.Logger.Critical(err)
		return err
	}

	return nil
}

func setupLogger() {
	logFileHandler, err := os.OpenFile(housekeeper.SharedInformation.Configuration.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	housekeeper.SharedInformation.Logger = logging.MustGetLogger("logger")
	logging.SetFormatter(logging.MustStringFormatter(`%{color}%{shortfunc} ▶ %{level:.4s} %{color:reset} %{message}`))

	backend1 := logging.NewLogBackend(os.Stdout, "", 0)

	logFile := logging.NewLogBackend(logFileHandler, "", 0)
	logFileLeveled := logging.AddModuleLevel(logFile)
	logFileLeveled.SetLevel(logging.ERROR, "")

	logging.SetBackend(backend1, logFileLeveled)

	if err != nil {
		housekeeper.SharedInformation.Logger.Critical(err)
		os.Exit(4)
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	environmentVariableConfigPath := os.Getenv("HOUSEKEEPER_CONFIGURATION_PATH")

	if len(environmentVariableConfigPath) > 0 {
		CONFIGURATION_PATH = environmentVariableConfigPath
	}

	err := loadConfiguration()

	setupLogger()

	if err != nil {
		housekeeper.SharedInformation.Logger.Critical(err)
		os.Exit(3)
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func(c chan os.Signal) {
		sig := <-c
		housekeeper.SharedInformation.Logger.Infof("Caught signal %s: shutting down.", sig)
		os.Exit(0)
	}(sigc)

	err = housekeeper.ConnectMQTTClient()

	if err != nil {
		housekeeper.SharedInformation.Logger.Error(err)
		housekeeper.SharedInformation.Logger.Error("Failed to connect to MQTT broker.")
	}

	err = housekeeper.StartWebserver()

	if err != nil {
		housekeeper.SharedInformation.Logger.Critical(err)
		os.Exit(2)
	}

	for true {
		time.Sleep(time.Second * 5)
	}
}
