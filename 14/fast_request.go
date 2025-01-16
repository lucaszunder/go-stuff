package main

import (
	"net/http"
	"time"
)

func FastRequest(uriA, uriB string) (fastest string) {
	durationA := MeasureTime(uriA)
	durationB := MeasureTime(uriB)

	if durationA < durationB {
		return uriA
	}

	return uriB
}

func MeasureTime(uri string) time.Duration {
	start := time.Now()
	http.Get(uri)
	return time.Since(start)
}
