package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFastRequest(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

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
