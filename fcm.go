package fcm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const fcmURL = "https://fcm.googleapis.com/fcm/send"

// Notification defines body for notification
type Notification struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}

// SendBody provides sending of the request to Google FCM
// More details on https://firebase.google.com/docs/cloud-messaging/http-server-ref
type SendBody struct {
	Notification Notification `json:"notification,omitempty"`
	// To represents a fcm token to device
	To string `json:"to,omitempty"`

	//Message with payload — data message
	Data interface{} `json:"data,omitempty"`

	Condition             string `json:"condition,omitempty"`
	CollapseKey           string `json:"collapse_key,omitempty"`
	Priority              string `json:"priority,omitempty"`
	ContentAvailable      bool   `json:"content_available,omitempty"`
	MutableContent        string `json:"mutable_content,omitempty"`
	TimeToLive            int    `json:"time_to_live,omitempty"`
	RestrictedPackageName string `json:"restricted_package_name,omitempty"`
	DryRun                bool   `json:"dry_run,omitempty"`
	TestURL               string `json:"-"`
}

// Response provides output data after sending to
// google FCM. More details on https://firebase.google.com/docs/cloud-messaging/http-server-ref#interpret-downstream
type Response struct {
	MulticastID int64    `json:"multicast_id"`
	Success     int      `json:"success"`
	Failure     int      `json:"failure"`
	Results     []Result `json:"results"`
}

// Result defines result data after response
type Result struct {
	MessageID      string `json:"message_id"`
	RegistrationID string `json:"registration_id"`
	Error          string `json:"error"`
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
func (a *App) Send(s *SendBody) (*Response, error) {
	marshalled, err := json.Marshal(s)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal body: %v", err)
	}

	url := fcmURL
	if s.TestURL != "" {
		url = s.TestURL + "/fcm/send"
		fmt.Println(url)
	}
	resp, err := a.sendRequest(url, marshalled)
	if err != nil {
		return nil, fmt.Errorf("unable to send request: %v", err)
	}
	defer func() {
		errC := resp.Close()
		if errC != nil {
			panic(fmt.Errorf("unable to close response body: %v", errC))
		}
	}()
	return a.getResponse(resp)
}

// sending of POST HTTP request
func (a *App) sendRequest(url string, b []byte) (io.ReadCloser, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("key=%s", a.serverKey))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	result := resp.Body
	return result, nil
}

// getResponse returns response after sending request to FCM API
func (a *App) getResponse(resp io.ReadCloser) (*Response, error) {
	var result *Response
	decoder := json.NewDecoder(resp)
	err := decoder.Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("unable to decode response: %v", err)
	}
	return result, nil
}
