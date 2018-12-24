package fcm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const fcmURL = "https://fcm.googleapis.com/fcm/send"

// Notification defines body for notification
type Notification struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}

// SendBody provides sending of the request to Google FCM
type SendBody struct {
	Notification Notification `json:"notification,omitempty"`
	// To represents a fcm token to device
	To string `json:"to,omitempty"`

	//Message with payload — data message
	Data interface{} `json:"data,omitempty"`
}

// App defines main definition for the app
type App struct {
	serverKey string
}

// New provides creating of the new app
func New(serverKey string) *App {
	return &App{
		serverKey: serverKey,
	}
}

// Send provides creating of the message to FCM
func (a *App) Send(s *SendBody) error {
	marshalled, err := json.Marshal(s)
	if err != nil {
		return fmt.Errorf("unable to marshal body: %v", err)
	}

	resp, err := a.sendRequest(marshalled)
	if err != nil {
		return fmt.Errorf("unable to send request: %v", err)
	}
	defer func() {
		errC := resp.Close()
		if errC != nil {
			panic(fmt.Errorf("unable to close response body: %v", errC))
		}
	}()
	body, err := ioutil.ReadAll(resp)
	if err != nil {
		return fmt.Errorf("unable to read body: %v", err)
	}

	return nil
}

func (a *App) sendRequest(b []byte) (io.ReadCloser, error) {
	req, err := http.NewRequest("POST", fcmURL, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("key=%s", "AAAAeOL9fKw:APA91bGTKBt7u_ftzsfWtp6yQKphudPlryQ1ufsLh09Q1UPNg0R5tuHJDNsCOYlaktxTurM03ufSFZNzIFJLLB6SjQd3A0SUlBTqlamwsLJEcRhj09Q7VPiZx6PL9pdhIXTCyB1jOXhq"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	result := resp.Body
	return result, nil
}
