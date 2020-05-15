package main

import(
	"github.com/op/go-logging"
	"github.com/alecthomas/kingpin"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"io/ioutil"
	//"net/http"
	//"net"
	"encoding/json"
	"./housekeeper"
	//"fmt"
	"os"
	"plugin"
//	"strings"
	"time"
	"math/rand"
//	"./cloudflare"
//	"./cfdnsu"
	"syscall"
	"os/signal"
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
	PLUGIN_PATH string
)

var (
	kingpinApp = kingpin.New("CFDNSU", "Cloudflare DNS updater")
	kingpinDump = kingpinApp.Command("dump", "Dump zone_identifiers and identifiers")
	kingpinRun = kingpinApp.Command("run", "Run CFDNSU in foreground").Default()

	log = logging.MustGetLogger("logger")
	pluginMap map[string]*plugin.Plugin
	eventMap map[string][]plugin.Symbol
	configuration s_configuration
)

type s_configuration struct {
	MQTT struct {
		Broker string `json:"broker"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"mqtt"`
	Plugin struct {
		Load []string `json:"load"`
		Configuration json.RawMessage `json:"configuration"`
	} `json:"plugin"`
}

func loadConfiguration() error {
	configurationData, err := ioutil.ReadFile(CONFIGURATION_PATH)

/*	log.Critical(configurationData)
	log.Critical(CONFIGURATION_PATH)
	log.Critical("awat")*/

	if err != nil {
		log.Critical(err)
		return err
	}

	err = json.Unmarshal(configurationData, &configuration)

	if err != nil {
		log.Critical(err)
		return err
	}

	housekeeper.SharedInformation.Configuration = configuration.Plugin.Configuration

//	cloudflare.Auth = configuration.Auth
//	cloudflare.Records = configuration.Records

	return nil

}

func loadPlugins() {
	var fullPath string
	var symbol plugin.Symbol

	eventMap = map[string][]plugin.Symbol{}

	if len(configuration.Plugin.Load) > 0 {
		for _, record := range configuration.Plugin.Load {
			fullPath = PLUGIN_PATH + "/" + record

			hotPlug, err := plugin.Open(fullPath)

			if err != nil {
				log.Critical(fullPath, err)
				continue
			}

			for _, eventName := range []string{"Startup", "Shutdown", "IpChanged", "IpUpdated"} {
				symbol, err = hotPlug.Lookup(eventName)

				if err == nil {
					eventMap[eventName] = append(eventMap[eventName], symbol)
				}
			}
		}
	}
}

func triggerEvent(eventName string) {
	log.Infof("Event triggered %s", eventName)
	if val, ok := eventMap[eventName]; ok {
		for _, element := range val {
			err := element.(func() error)()

			if err != nil {
				log.Error(err)
			}
		}
	}
}

/*func resolveIp() (error, string) {
	url := configuration.Check.Targets[rand.Intn(len(configuration.Check.Targets))]
	resp, err := http.Get(url)

	if err != nil {
		return err, ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err, ""
	}

	if resp.StatusCode > 200 {
		return fmt.Errorf("Wrong response code %d", resp.StatusCode), ""
	}

	ip := strings.TrimSpace(string(body))

	if net.ParseIP(ip).To4() == nil {
		return fmt.Errorf("Not ipv4 coming from %s", url), ""
	}

	return nil, ip
}*/



/*func dump() {
	err, cFListZones := cloudflare.GetCFListZones()

	if err != nil {
		log.Critical(err)

		return
	}

	if !cFListZones.Success {
		log.Critical(cFListZones.Errors)

		return
	}

	for _, zone := range cFListZones.Result {
		err, zoneRecord := cloudflare.GetCFListDNSRecords(zone.Id)

		if err != nil {
			log.Error(err)
		}

		if zoneRecord.Success == false {
			log.Errorf("%+v", zone)
		}

		zoneMax := len(zoneRecord.Result) - 1

		modifier := "┌"

		if zoneMax == -1 {
			modifier = " "
		}

		log.Infof("%s[%s][%s]", modifier, zone.Id, zone.Name)

		for iterator, element := range zoneRecord.Result {
			err, zoneDetails := cloudflare.GetCFDNSRecordDetails(zone.Id, element.Id)

			if err != nil {
				log.Error(err)
			}

			modifier = "├"

			if iterator == zoneMax {
				modifier = "└"
			}

			log.Infof("%s─%s - %s", modifier, zoneDetails.Result.Id, zoneDetails.Result.Name)
		}
	}
}*/

func run() {
	housekeeper.SharedInformation.Logger = log

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func(c chan os.Signal) {
		sig := <-c
		log.Infof("Caught signal %s: shutting down.", sig)
		os.Exit(0)
	}(sigc)

//	oldIp := ""

	go triggerEvent("Startup")

	for true {
		time.Sleep(time.Second * 5)
	}

/*	for true {
		err, currentIp := resolveIp()

		if err != nil {
			if oldIp == "" {
				log.Fatal(err)
				return
			} else {
				log.Warning(err)
			}
		} else {
			if currentIp != oldIp {
				cfdnsu.SharedInformation.CurrentIp = currentIp
				log.Infof("Current ip %s", currentIp)
				go triggerEvent("IpChanged")

				for recordIndex, record := range cloudflare.Records {
					err, cFDNSRecordDetails := cloudflare.GetCFDNSRecordDetails(record.ZoneIdentifier, record.Identifier)

					if err != nil {
						log.Error(err)
						continue
					}

					if !cFDNSRecordDetails.Success {
						log.Error(cFDNSRecordDetails.Errors)
						continue
					}

					cloudflare.Records[recordIndex].Name = cFDNSRecordDetails.Result.Name

					if cFDNSRecordDetails.Result.Ip != currentIp {
						err, cCFDNSRecord := cloudflare.SetCFDNSRecord(recordIndex, currentIp)

						if err != nil {
							log.Error(err)
							log.Error(cCFDNSRecord)
							log.Error("Failed to update record")
							continue
						}

						if !cCFDNSRecord.Success {
							log.Error(cCFDNSRecord)
							continue
						}

						go triggerEvent("IpUpdated")

						log.Infof("Server ip has changed to %s previously %s updating, updated %t", currentIp, cFDNSRecordDetails.Result.Ip, cCFDNSRecord.Success)
					}
				}
			}

			oldIp = currentIp
		}

		time.Sleep(time.Second * time.Duration(configuration.Check.Rate))
	}*/
}

func main() {
	var err error
	rand.Seed(time.Now().Unix())
	logging.SetFormatter(logging.MustStringFormatter(`%{color} %{shortfunc} ▶ %{level:.4s} %{color:reset} %{message}`))

	var environmentVariableConfigPath = os.Getenv("CFDNSU_CONFIGURATION_PATH")
	var environmentVariablePluginPath = os.Getenv("CFDNSU_PLUGIN_PATH")

	if len(environmentVariableConfigPath) > 0 {
		CONFIGURATION_PATH = environmentVariableConfigPath
	}

	if len(environmentVariablePluginPath) > 0 {
		PLUGIN_PATH = environmentVariablePluginPath
	}

	err = loadConfiguration()
	loadPlugins()

	if err != nil {
		log.Criticalf("%s", err)
		return
	}

//	log.Criticalf("%s", configuration)
//	log.Criticalf("%s", configuration.MQTT.Broker)

	mqttOptions := MQTT.NewClientOptions()

	mqttOptions.AddBroker(configuration.MQTT.Broker)
	mqttOptions.SetUsername(configuration.MQTT.Username)
	mqttOptions.SetPassword(configuration.MQTT.Password)

/*	log.Criticalf("%s", configuration.MQTT.Broker)
	log.Criticalf("%s", configuration.MQTT.Username)
	log.Criticalf("%s", configuration.MQTT.Password)*/

	client := MQTT.NewClient(mqttOptions)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	housekeeper.SharedInformation.MQTTClient = client

/*	client.Subscribe("#", 0, func(client MQTT.Client, msg MQTT.Message) {
		log.Criticalf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})*/

	run()

/*	switch kingpin.MustParse(kingpinApp.Parse(os.Args[1:])) {
		case kingpinDump.FullCommand():
			dump()
		case kingpinRun.FullCommand():
			run()
	}*/
}
