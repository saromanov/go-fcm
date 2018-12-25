# Go-fcm

Implementation of Google Cloud Messaging API

## Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/saromanov/go-fcm"
)

func main() {
	data := fcm.New("YOUR_SERVER_KEY")
	resp, err := data.Send(&fcm.SendBody{
		To: "TOKEN",
		Notification: fcm.Notification{
			Title: "TestText",
			Body:  "TestBody",
		},
		Data: map[string]string{"body": "TestBody", "title": "TestText"},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

```