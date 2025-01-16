package main

import (
	"errors"
	"net/http"
	"time"
)

func FastRequest(uriA, uriB string, timeout time.Duration) (string, error) {
	select {
	case <-Ping(uriA):
		return uriA, nil
	case <-Ping(uriB):
		return uriB, nil
	case <-time.After(timeout):
		return "", errors.New("timeout")
	}
}

func Ping(uri string) chan bool {
	response := make(chan bool)

	go func(channel chan bool) {
		http.Get(uri)
		response <- true
	}(response)
	return response
}

func MeasureTime(uri string) time.Duration {
	start := time.Now()
	http.Get(uri)
	return time.Since(start)
}
