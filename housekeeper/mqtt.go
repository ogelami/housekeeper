package housekeeper

import(
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func ConnectMQTTClient() error {
	mqttOptions := MQTT.NewClientOptions()

	mqttOptions.AddBroker(SharedInformation.Configuration.MQTT.Broker)
	mqttOptions.SetUsername(SharedInformation.Configuration.MQTT.Username)
	mqttOptions.SetPassword(SharedInformation.Configuration.MQTT.Password)

	/*	log.Criticalf("%s", configuration.MQTT.Broker)
	log.Criticalf("%s", configuration.MQTT.Username)
	log.Criticalf("%s", configuration.MQTT.Password)*/

	client := MQTT.NewClient(mqttOptions)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	SharedInformation.Logger.Info("Connected")

	SharedInformation.MQTTClient = client

	return nil
}