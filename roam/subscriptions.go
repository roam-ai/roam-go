package roam

import (
	"roam-go/roam/helpers"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type subscriptionType string

var (
	projectSubscription subscriptionType = "Project Subscription"
	userSubscription    subscriptionType = "User Subscription"
	groupSubscription   subscriptionType = "Group Subscription"
)

// Subscription is interface definition for a roam-go client
type Subscription interface {
	// Subscribe will subscribe and trigger handler func on message
	Subscribe(handler MessageHandler) error
	//Unsubscribe will unsubscribe from subscibed topics
	Unsubscribe() error
}

// subscription is struct with variables needed to connect to Roam MQTT broker
type subscription struct {
	apikey     string
	accountID  string
	projectID  string
	mqttClient mqtt.Client
	groupID    string
	userID     string
	subType    subscriptionType
	groupUsers []string
}

// newSubscription implements Client interface
func newSubscription(apiKey string) (*subscription, error) {
	accountID, projectID, err := helpers.GetProjectDetails(apiKey)
	if err != nil {
		return nil, err
	}
	c := &subscription{
		apikey:    apiKey,
		accountID: accountID,
		projectID: projectID,
	}
	return c, nil
}

// NewProjectSubscription to initialize subscription
// to subscribe to project level data
func NewProjectSubscription(apikey string) (Subscription, error) {
	c, err := newSubscription(apikey)
	if err != nil {
		return nil, err
	}
	c.subType = projectSubscription
	return c, err
}

// NewGroupSubscription to initialize subscription
// to subscribe to group level data
func NewGroupSubscription(apikey, groupID string) (Subscription, error) {
	c, err := newSubscription(apikey)
	if err != nil {
		return nil, err
	}
	c.groupID = groupID
	c.subType = groupSubscription
	return c, nil
}

// NewUserSubscription to initialize subscription
// to subscribe to a single user data
func NewUserSubscription(apikey, userID string) (Subscription, error) {
	c, err := newSubscription(apikey)
	if err != nil {
		return nil, err
	}
	c.userID = userID
	c.subType = userSubscription
	return c, nil
}
