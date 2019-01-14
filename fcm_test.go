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
		Test: true,
		To:   "test",
		Data: map[string]string{
			"foo": "bar",
		},
	})
	assert.NoError(t, err, "unable to send data")

}

func mockFCMServer() *httptest.Server {
	router := chi.NewRouter()
	router.Post("/fcm/send", func(w http.ResponseWriter, r *http.Request) {
		s := `{}`
		_, _ = w.Write([]byte(s))
	})
	return httptest.NewServer(router)
}
