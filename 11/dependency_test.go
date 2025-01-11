package main

import (
	"bytes"
	"testing"
)

func TestDependency(t *testing.T) {
	buffer := bytes.Buffer{}
	Greetings(&buffer, "Chris")

	output := buffer.String()
	expectation := "Hello, Chris"

	if output != expectation {
		t.Errorf("Output '%s', expectation '%s'", output, expectation)
	}
}
