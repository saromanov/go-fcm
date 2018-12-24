package fcm

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
