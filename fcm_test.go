package fcm

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	a = New("testKey")
	assert.NotNil(t, a, "unable to create object")
}

func mockFCMServer() *httptest.Server {
	router := chi.NewRouter()
	router.Get("/fcm/send", func(w http.ResponseWriter, r *http.Request) {
		s := `{}`
		_, _ = w.Write([]byte(s))
	})
	return httptest.NewServer(router)
}
