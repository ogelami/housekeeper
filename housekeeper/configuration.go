package housekeeper

import(
	"github.com/op/go-logging"
	"encoding/json"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type sharedInformation struct {
	Logger *logging.Logger
	Configuration json.RawMessage
	MQTTClient MQTT.Client
}

var SharedInformation = sharedInformation{ nil, nil, nil }

func ParseConfig (v interface{}) error {
	err := json.Unmarshal(SharedInformation.Configuration, &v)
	
	if err != nil {
		SharedInformation.Logger.Error(err)
	}

	return err
}
