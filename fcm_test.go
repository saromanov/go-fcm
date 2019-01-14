package fcm

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	a := New("testKey")
	assert.NotNil(t, a, "unable to create object")
}

func TestSend(t *testing.T) {
	s := mockFCMServer()
	defer s.Close()
	a := New("testKey")
	assert.NotNil(t, a, "unable to create object")
	_, err := a.Send(&SendBody{
		TestURL: s.URL + "/fcm/send",
		To:      "test",
		Data: map[string]string{
			"foo": "bar",
		},
	})
	assert.NoError(t, err, "unable to send data")
	_, err = a.Send(&SendBody{
		TestURL: s.URL + "/badjson",
		To:      "test",
		Data: map[string]string{
			"foo": "bar",
		},
	})
	assert.Error(t, err, "unable to decode response: json: cannot unmarshal string into Go value of type fcm.Response")

	_, err = a.Send(&SendBody{
		TestURL: "127.0.2.2",
		To:      "test",
		Data: map[string]string{
			"foo": "bar",
		},
	})
	assert.Error(t, err)

}

func mockFCMServer() *httptest.Server {
	router := chi.NewRouter()
	router.Post("/fcm/send", func(w http.ResponseWriter, r *http.Request) {
		s := `{
			"multicast_id":1,
			"success":1,
			"failure":0,
			"results":[]
		}`
		_, _ = w.Write([]byte(s))
	})
	router.Post("/badjson", func(w http.ResponseWriter, r *http.Request) {
		s := `
			"multicast_id":1,
			"success":1,
			"failure":0,
			"results":[]
		}`
		_, _ = w.Write([]byte(s))
	})
	return httptest.NewServer(router)
}
