# Go-fcm
[![Go Report Card](https://goreportcard.com/badge/github.com/saromanov/go-fcm)](https://goreportcard.com/report/github.com/saromanov/go-fcm)
[![Build Status](https://travis-ci.org/saromanov/go-fcm.svg?branch=master)](https://travis-ci.org/saromanov/go-fcm)
[![Coverage Status](https://coveralls.io/repos/github/saromanov/go-fcm/badge.svg?branch=master)](https://coveralls.io/github/saromanov/go-fcm?branch=master)
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