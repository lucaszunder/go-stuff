package main

import "testing"

func TestSumArray(t *testing.T) {
	output := SumArray([3]int{1, 2, 3})
	expectation := 6

	if output != expectation {
		t.Errorf("The output for the method Sum is '%v' and the expectation was '%v'", output, expectation)
	}
}
