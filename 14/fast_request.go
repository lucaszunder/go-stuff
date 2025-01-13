package main

import (
	"net/http"
	"time"
)

func FastRequest(uriA, uriB string) (fastest string) {
	startA := time.Now()
	http.Get(uriA)
	durationA := time.Since(startA)

	startB := time.Now()
	http.Get(uriB)
	durationB := time.Since(startB)

	if durationA < durationB {
		return uriA
	}

	return uriB
}
