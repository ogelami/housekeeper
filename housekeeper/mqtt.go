package housekeeper

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type S_MQTTResponse struct {
	Topic   string `json:"topic"`
	Message string `json:"message"`
}

func PublishMQTTMessage(topic string, payload string) {
	client.Publish(topic, 0, true, payload)
}

var client MQTT.Client

func ConnectMQTTClient() error {
	mqttOptions := MQTT.NewClientOptions()

	mqttOptions.AddBroker(Configuration.MQTT.Broker)
	mqttOptions.SetUsername(Configuration.MQTT.Username)
	mqttOptions.SetPassword(Configuration.MQTT.Password)

	/*	log.Criticalf("%s", configuration.MQTT.Broker)
		log.Criticalf("%s", configuration.MQTT.Username)
		log.Criticalf("%s", configuration.MQTT.Password)*/

	client = MQTT.NewClient(mqttOptions)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	Logger.Info("Connected")

	client.Subscribe("#", 0, func(client MQTT.Client, msg MQTT.Message) {
		serverResponse := &S_MQTTResponse{Topic: string(msg.Topic()), Message: string(msg.Payload())}

		//		Logger.Info(serverResponse)

		Hub.broadcast <- serverResponse
	})

	return nil
}
