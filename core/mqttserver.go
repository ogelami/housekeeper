package core

import (
	"log"

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

func serve(server *mqtt.Server) error {
	err := serverHandler.Serve()

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func StartMQTTserver() error {
	serverHandler := mqtt.New()

	tcp := listeners.NewTCP("t1", Configuration.MQTT.Listen)

	err := serverHandler.AddListener(tcp, nil)

	if err != nil {
		log.Fatal(err)
	}

	serverHandler.Events.OnMessage = func(cl events.Client, pk events.Packet) (pkx events.Packet, err error) {
		serverResponse := &S_MQTTResponse{Topic: string(pk.TopicName), Message: string(pk.Payload)}

		Hub.broadcast <- serverResponse

		return pk, nil
	}

	Logger.Infof("Serving MQTT server")

	go serve(serverHandler)

	return nil
}
