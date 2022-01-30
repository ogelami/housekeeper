package internal

import (
	"encoding/json"
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
	Logger.Info(serverHandler)

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

func StartMQTTserver() error {
	serverHandler = mqtt.New()

	tcp := listeners.NewTCP("t1", Configuration.MQTT.Listen)
	Logger.Info(tcp)
	err := serverHandler.AddListener(tcp, nil)

	if err != nil {
		Logger.Fatal(err)
	}

	/*serverHandler.Events.OnConnect = func(client events.Client, packet events.Packet) {
		Logger.Info("C con")
	}*/

	serverHandler.Events.OnMessage = func(client events.Client, packet events.Packet) (pkx events.Packet, err error) {
		serverResponse := &S_MQTTResponse{Topic: packet.TopicName, Message: string(packet.Payload)}

		if strings.HasPrefix(packet.TopicName, "tasmota/discovery") && strings.HasSuffix(packet.TopicName, "/config") {
			k := DiscoveryMessage{}

			err := json.Unmarshal(packet.Payload, &k)

			if err != nil {
				Logger.Warning(err)
			}

			DeviceMap[k.Topic] = *serverResponse

			Logger.Info(k)
		} else if strings.HasSuffix(packet.TopicName, "/LWT") {
			delete(DeviceMap, "") // <--
		}

		Hub.broadcast <- serverResponse

		return packet, err
	}

	go serve()

	Logger.Infof("Serving MQTT server, listening on %s", Configuration.MQTT.Listen)

	return nil
}
