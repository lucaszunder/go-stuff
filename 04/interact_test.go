package main

import "testing"

func TestInteract(t *testing.T) {
	output := Interact("A")
	expectation := "AAAAA"

	if output != expectation {
		t.Errorf("The output for the method Sum is '%s' and the expectation was '%s'", output, expectation)
	}
}

func BenchmarkInteract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Interact("A")
	}
}
