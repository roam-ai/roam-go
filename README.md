# roam-go
A Golang backend library for GeoSpark Location Subscription. It is used to subscribe to user's locations at project level or location of a single user.

roam-go supports subscription to the following location data:
* Specific user
* All users of a group
* All users of project

## Installation
You can install our Golang library as mentioned below.
```
go get -u github.com/roam-ai/roam-go/roam
```

## Example Usage
You can think this library as a wrapper around our REST API which needs your API key for authorization and it works as per project level. It is fairly simple to use:
- Create an instance of subscription with your API key. 
- Define your custom callback function and pass it with Subscribe method, which will be executed on every location received from the server.
- To stop receiving data, call Unsubscribe method.
The below example usage code:

```
package main

import (
	"fmt"
	"github.com/roam-ai/roam-go/roam"
	"time"
)
func main() {
    // initialization
    subscription , _ := roam.NewProjectSubscription("apikey")

    var handler roam.MessageHandler = func(userID string , payload []byte){
        fmt.Println(userID)
        fmt.Println(string(payload))
	}

    subscription.Subscribe(handler)
}
```

### Subscribe to a single user
You can use the SDK to subscribe to single user's location and listen to it. You have pass the user id which our Mobile SDK's returns during creating a user.

```
userSubscription , err := roam.NewUserSubscription("apikey","userID")
```
#### Example Usage
```
package main

import (
	"fmt"
	"github.com/roam-ai/roam-go/roam"
	"time"
)
func main() {
    // initialization
    subscription , _ := roam.NewUserSubscription("apikey","userID")

    // Declaring handler function
    var handler roam.MessageHandler = func(userID string , payload []byte){
        // logic
	}

    // start subscription
    subscription.Subscribe(handler)
}
```

### Subscribe to a group:
You can also use the SDK to subscribe to a user group and listen to their location updates. 
You can get the group_id from our developer API for user grouping and use the group id when creating the client instance to listen to location at user group level.

```
groupSubscription , err := roam.NewGroupSubscription("apikey","groupID")
```
```
package main

import (
	"fmt"
	"github.com/roam-ai/roam-go/roam"
	"time"
)
func main() {
    // initialization
    subscription , _ := roam.NewGroupSubscription("apikey","groupID")

    // Declaring handler function
    var handler roam.MessageHandler = func(userID string , payload []byte){
        // logic
	}

    // start subscription
    subscription.Subscribe(handler)
}
```

### Unsubscribe
At any point, if you wish to unsubscribe from location updates, use the below code:
```
subscription.Unsubscribe()
```

### Handler Function
While calling Subscribe method on subscription, you will need to pass a handler func as parameter.
Handler func should have userID of type string and payload of type []byte and will be invoked on every location update received.
```
    var handler roam.MessageHandler = func(userID string , payload []byte){
        // logic
	}

    subscription.Subscribe(handler)
```
