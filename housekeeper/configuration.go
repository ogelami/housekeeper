package housekeeper

import(
	"github.com/op/go-logging"
	"encoding/json"
)

type sharedInformation struct {
	Logger *logging.Logger
	Configuration json.RawMessage
}

var SharedInformation = sharedInformation{ nil, nil }

func ParseConfig (v interface{}) error {
	err := json.Unmarshal(SharedInformation.Configuration, &v)
	
	if err != nil {
		SharedInformation.Logger.Error(err)
	}

	return err
}
