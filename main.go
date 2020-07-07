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

	/*	housekeeper.Logger.Critical(configurationData)
		housekeeper.Logger.Critical(CONFIGURATION_PATH)*/

	if err != nil {
		housekeeper.Logger.Critical(err)
		return err
	}

	err = json.Unmarshal(configurationData, &housekeeper.Configuration)

	//	housekeeper.Logger.Critical(configuration)

	if err != nil {
		housekeeper.Logger.Critical(err)
		return err
	}

	return nil
}

func setupLogger() {
	housekeeper.Logger = logging.MustGetLogger("logger")
	logging.SetFormatter(logging.MustStringFormatter(`%{color}%{shortfunc} ▶ %{level:.4s} %{color:reset} %{message}`))

	backend1 := logging.NewLogBackend(os.Stdout, "", 0)

	if len(housekeeper.Configuration.LogFile) > 0 {
		logFileHandler, err := os.OpenFile(housekeeper.Configuration.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			housekeeper.Logger.Critical(err)
			os.Exit(4)
		}

		logFile := logging.NewLogBackend(logFileHandler, "", 0)
		backend2 := logging.AddModuleLevel(logFile)
		backend2.SetLevel(logging.ERROR, "")

		logging.SetBackend(backend1, backend2)
	} else {
		logging.SetBackend(backend1)
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
		housekeeper.Logger.Critical(err)
		os.Exit(3)
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func(c chan os.Signal) {
		sig := <-c
		housekeeper.Logger.Infof("Caught signal %s: shutting down.", sig)
		os.Exit(0)
	}(sigc)

	err = housekeeper.ConnectMQTTClient()

	if err != nil {
		housekeeper.Logger.Error(err)
		housekeeper.Logger.Error("Failed to connect to MQTT broker.")
	}

	err = housekeeper.StartWebserver()

	if err != nil {
		housekeeper.Logger.Critical(err)
		os.Exit(2)
	}

	for true {
		time.Sleep(time.Second * 5)
	}
}
