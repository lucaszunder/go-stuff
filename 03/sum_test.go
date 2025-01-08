package main

import "testing"

func TestSoma(t *testing.T) {
	result := Sum(1, 2)
	expectation := 3

	if result != expectation {
		t.Errorf("The output for the method Sum is '%v' and the expectation was '%v'", result, expectation)
	}
}
