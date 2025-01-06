package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Teste")
	expectation := "Ola Teste"

	if result != expectation {
		t.Errorf("result '%s', expected '%s'", result, expectation)
	}
}
