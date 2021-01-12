# roam-go
Roam Golang SDK for subscription to real-time location data.

roam-go supports subscription to the following location data:
* Specific user
* All users of a group
* All users of project

## Installation
```
go get -u github.com/geosparks/roam-go/roam
```

## Example
### Subscribe to a single user:
```
package main

import (
	"fmt"
	"github.com/geosparks/roam-go/roam"
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

    // stop subscription
    subscription.Unsubscribe()
}
```

### Subscribe to a group:
```
package main

import (
	"fmt"
	"github.com/geosparks/roam-go/roam"
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

    // stop subscription
    subscription.Unsubscribe()
}
```

### Subscribe to a whole project:
```
package main

import (
	"fmt"
	"github.com/geosparks/roam-go/roam"
	"time"
)
func main() {
    // initialization
    subscription , _ := roam.NewProjectSubscription("apikey")

    // Declaring handler function
    var handler roam.MessageHandler = func(userID string , payload []byte){
        // logic
	}

    // start subscription
    subscription.Subscribe(handler)

    // stop subscription
    subscription.Unsubscribe()
}
```