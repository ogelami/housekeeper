package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

	housekeeper "github.com/ogelami/housekeeper/internal"
	logging "github.com/op/go-logging"
)

var (
	CONFIGURATION_PATH string = "housekeeper.conf"
	SYSCONF_PATH       string = ""
)

func setupLogger() {
	housekeeper.Logger = logging.MustGetLogger("logger")

	logging.SetFormatter(logging.MustStringFormatter(`%{color}%{shortfunc} ▶ %{level:.4s} %{color:reset} %{message}`))
	logging.SetBackend(logging.NewLogBackend(os.Stdout, "", 0))
}

func loadConfiguration() error {
	var configurationData []byte
	var err error

	configurationPath, configurationPathSet := os.LookupEnv("HOUSEKEEPER_CONFIGURATION_PATH")

	if configurationPathSet {
		configurationData, err = ioutil.ReadFile(configurationPath)
	} else if len(SYSCONF_PATH) > 0 {
		configurationData, err = ioutil.ReadFile(SYSCONF_PATH + "/" + CONFIGURATION_PATH)
	} else {
		configurationData, err = ioutil.ReadFile(CONFIGURATION_PATH)
	}

	if err != nil {
		housekeeper.Logger.Critical(err)
		return err
	}

	err = json.Unmarshal(configurationData, &housekeeper.Configuration)

	if err != nil {
		housekeeper.Logger.Critical(err)
		return err
	}

	return nil
}

func main() {
	setupLogger()
	err := loadConfiguration()

	if err != nil {
		housekeeper.Logger.Critical(err)
		os.Exit(3)
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)

	go func(c chan os.Signal) {
		sig := <-c
		housekeeper.Logger.Infof("Caught signal %s: shutting down.", sig)
		os.Exit(0)
	}(sigc)

	err = housekeeper.StartMQTTserver()

	if err != nil {
		housekeeper.Logger.Error(err)
		housekeeper.Logger.Error("Failed to start MQTT broker.")
	}

	err = housekeeper.StartWebserver()

	if err != nil {
		housekeeper.Logger.Critical(err)
		os.Exit(2)
	}

	for {
		time.Sleep(time.Second * 5)
	}
}
