package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFastRequest(t *testing.T) {
	slowServer := CreateFakeHttpServer(20 * time.Millisecond)

	fastServer := CreateFakeHttpServer(5 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURI := slowServer.URL
	fastURI := fastServer.URL

	expectation := fastURI
	output := FastRequest(slowURI, fastURI)

	if output != expectation {
		t.Errorf("resultado '%s', esperado '%s'", output, expectation)
	}
}

func CreateFakeHttpServer(sleepTime time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(sleepTime)
		w.WriteHeader(http.StatusOK)
	}))
}
