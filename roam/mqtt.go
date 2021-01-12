package roam

import (
	"crypto/sha512"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

func generateMQTTParams(apikey string) (clientID string, username string, password string) {
	timestamp := time.Now().UTC().Unix()
	uniqueID := uuid.New().String()
	clientID = fmt.Sprintf("%s_%s", apikey, uniqueID)
	username = fmt.Sprintf("api_%d?x-amz-customauthorizer-name=%s", timestamp, authorizerName)
	sha := sha512.New()
	sha.Write([]byte(fmt.Sprintf("%s%d%s", apikey, timestamp, salt)))

	password = hex.EncodeToString(sha.Sum(nil))
	return clientID, username, password
}

func generateMQTTClient(apikey string) (client mqtt.Client, err error) {
	clientID, username, password := generateMQTTParams(apikey)

	tlsConfig := &tls.Config{
		NextProtos: []string{"mqtt"},
	}
	// mqtt.DEBUG = log.New(os.Stdout, "Debug from Paho: ", log.LstdFlags)
	opts := mqtt.NewClientOptions()
	broker := fmt.Sprintf("tcps://%s:443", defaultIOTURL)
	opts.AddBroker(broker)
	opts.SetUsername(username).SetPassword(password)
	opts.SetClientID(clientID)
	opts.SetTLSConfig(tlsConfig)
	c := mqtt.NewClient(opts)
	token := c.Connect()
	if token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return c, nil
}
