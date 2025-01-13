package main

import "testing"

func TestFastRequest(t *testing.T) {
	slowURI := "http://www.facebook.com"
	fastURI := "http://www.quii.co.uk"

	expectation := fastURI
	output := FastRequest(slowURI, fastURI)

	if output != expectation {
		t.Errorf("resultado '%s', esperado '%s'", output, expectation)
	}
}
