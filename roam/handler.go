package roam

import (
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	defaultIOTURL  = "az91jf6dri5ey-ats.iot.eu-central-1.amazonaws.com"
	authorizerName = "iot-authorizer"
	prefix         = ""
	salt           = "x9nFgM1ioxAOPmT3Fdyeh483lerc1J7k"
)

// MessageHandler is handler function for roam-go
// MessageHandler func passed in with subscribe method
// will be invoked on receiving data from backend
type MessageHandler func(userID string, message []byte)

// internalMessageHandlerProxy will take the raw mqtt message
// and parse to get payload and userID and
// invoke provided handler function
func internalMessageHandlerProxy(handler MessageHandler) mqtt.MessageHandler {
	return func(client mqtt.Client, msg mqtt.Message) {
		topicSplit := strings.Split(msg.Topic(), "/")
		userID := topicSplit[len(topicSplit)-1]
		handler(userID, msg.Payload())
	}
}
