package roam

import (
	"errors"
	"fmt"

	"github.com/geosparks/roam-go/roam/helpers"
)

func (c *subscription) Subscribe(handler MessageHandler) error {
	mqttClient, err := generateMQTTClient(c.apikey)
	if err != nil {
		return errors.New("Connection error: " + err.Error())
	}
	c.mqttClient = mqttClient
	switch c.subType {
	case projectSubscription:
		token := c.mqttClient.Subscribe(fmt.Sprintf("%slocations/%s/%s/+", prefix, c.accountID, c.projectID), 0, internalMessageHandlerProxy(handler))
		if token.Wait() && token.Error() != nil {
			c.mqttClient.Disconnect(0)
			return token.Error()
		}
		return nil
	case userSubscription:
		token := c.mqttClient.Subscribe(fmt.Sprintf("%slocations/%s/%s/%s", prefix, c.accountID, c.projectID, c.userID), 0, internalMessageHandlerProxy(handler))
		if token.Wait() && token.Error() != nil {
			c.mqttClient.Disconnect(0)
			return token.Error()
		}
		return nil

	case groupSubscription:
		userIDs, err := helpers.GetGroupData(c.apikey, c.groupID)
		if err != nil {
			return err
		}
		c.groupUsers = userIDs
		for _, user := range userIDs {
			token := c.mqttClient.Subscribe(fmt.Sprintf("%slocations/%s/%s/%s", prefix, c.accountID, c.projectID, user), 0, internalMessageHandlerProxy(handler))
			if token.Wait() && token.Error() != nil {
				c.mqttClient.Disconnect(0)
				return token.Error()
			}
		}
		return nil
	default:
		c.mqttClient.Disconnect(0)
		return errors.New("Error with subscription. Please reinitiate")
	}

}

func (c *subscription) Unsubscribe() error {
	defer c.mqttClient.Disconnect(0)
	switch c.subType {
	case projectSubscription:
		token := c.mqttClient.Unsubscribe(fmt.Sprintf("%slocations/%s/%s/+", prefix, c.accountID, c.projectID))
		if token.Wait() && token.Error() != nil {
			return token.Error()
		}
		return nil
	case userSubscription:
		token := c.mqttClient.Unsubscribe(fmt.Sprintf("%slocations/%s/%s/%s", prefix, c.accountID, c.projectID, c.userID))
		if token.Wait() && token.Error() != nil {
			return token.Error()
		}
		return nil

	case groupSubscription:
		for _, user := range c.groupUsers {
			token := c.mqttClient.Unsubscribe(fmt.Sprintf("%slocations/%s/%s/%s", prefix, c.accountID, c.projectID, user))
			if token.Wait() && token.Error() != nil {
				return token.Error()
			}
		}
		return nil
	default:
		return errors.New("Error with subscription. Please reinitiate")
	}
}
