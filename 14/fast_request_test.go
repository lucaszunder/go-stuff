package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFastRequest(t *testing.T) {
	t.Run("Should select the most performatic url", func(t *testing.T) {
		slowServer := CreateFakeHttpServer(20 * time.Millisecond)

		fastServer := CreateFakeHttpServer(5 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURI := slowServer.URL
		fastURI := fastServer.URL

		timeout := 10 * time.Second

		expectation := fastURI
		output, _ := FastRequest(slowURI, fastURI, timeout)

		if output != expectation {
			t.Errorf("resultado '%s', esperado '%s'", output, expectation)
		}
	})

	t.Run("Should raise one error if the process took more then 10 seconds", func(t *testing.T) {
		slowServer := CreateFakeHttpServer(20 * time.Millisecond)

		fastServer := CreateFakeHttpServer(11 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURI := slowServer.URL
		fastURI := fastServer.URL

		timeout := 10 * time.Millisecond

		_, err := FastRequest(slowURI, fastURI, timeout)

		if err == nil {
			t.Errorf("Was expected one error")
		}
	})

}

func CreateFakeHttpServer(sleepTime time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(sleepTime)
		w.WriteHeader(http.StatusOK)
	}))
}
