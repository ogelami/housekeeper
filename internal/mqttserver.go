package internal

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"

	mqtt "github.com/mochi-co/mqtt/server"
	"github.com/mochi-co/mqtt/server/events"
	"github.com/mochi-co/mqtt/server/listeners"
)

type S_MQTTResponse struct {
	Topic   string `json:"topic"`
	Message string `json:"message"`
}

var serverHandler *mqtt.Server

func PublishMQTTMessage(topic string, payload string) {
	serverHandler.Publish(topic, []byte(payload), false)
}

func serve() error {
	err := serverHandler.Serve()

	if err != nil {
		Logger.Fatal(err)
	}

	return nil
}

type DiscoveryMessage struct {
	Mac        string   `json:"mac,omitempty"`
	Topic      string   `json:"t,omitempty"`
	Device     string   `json:"md,omitempty"`
	Ip         string   `json:"ip,omitempty"`
	State      []string `json:"state,omitempty"`
	Tp         []string `json:"tp,omitempty"`
	LWTOffline string   `json:"ofln,omitempty"`
	LWTOnline  string   `json:"onln,omitempty"`
}

var DeviceMap = make(map[string]S_MQTTResponse)
var LastPowerStatusMap = make(map[string]S_MQTTResponse)

func extractIdentifier(topic string) (string, error) {

	var rgx = regexp.MustCompile(`\/(.+?)\/`)
	rs := rgx.FindStringSubmatch(topic)

	if rs != nil {
		return rs[1], nil
	}

	return "", errors.New("identifier could not be found")
}

func StartMQTTserver() error {
	serverHandler = mqtt.New()

	tcp := listeners.NewTCP("t1", Configuration.MQTT.Listen)

	err := serverHandler.AddListener(tcp, nil)

	if err != nil {
		Logger.Fatal(err)
	}

	serverHandler.Events.OnMessage = func(client events.Client, packet events.Packet) (pkx events.Packet, err error) {
		serverResponse := &S_MQTTResponse{Topic: packet.TopicName, Message: string(packet.Payload)}

		if strings.HasPrefix(packet.TopicName, "tasmota/discovery") && strings.HasSuffix(packet.TopicName, "/config") {
			k := DiscoveryMessage{}

			err := json.Unmarshal(packet.Payload, &k)

			if err != nil {
				Logger.Warning(err)
			}

			DeviceMap[k.Topic] = *serverResponse

			Logger.Infof("Discovery %+v", k)
		} else if strings.HasPrefix(packet.TopicName, "stat") && strings.HasSuffix(packet.TopicName, "/POWER") {
			identifier, err := extractIdentifier(packet.TopicName)

			if err != nil {
				Logger.Warning(err)
				return packet, err
			}

			LastPowerStatusMap[identifier] = *serverResponse
		} else if strings.HasPrefix(packet.TopicName, "tele") && strings.HasSuffix(packet.TopicName, "/LWT") {
			identifier, err := extractIdentifier(packet.TopicName)

			if err != nil {
				Logger.Warning(err)
				return packet, err
			}

			Logger.Infof("LWT : %s - %+v", identifier, string(packet.Payload))
		}

		Hub.broadcast <- serverResponse

		return packet, err
	}

	go serve()

	Logger.Infof("Serving MQTT server, listening on %s", Configuration.MQTT.Listen)

	return nil
}
