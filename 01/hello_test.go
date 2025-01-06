package main

import "testing"

// requirements: if we call hello with a empty string, should return "Hello World", if we call it with a string
// should return "Hello STRING"

func TestHello(t *testing.T) {
	validateMessage := func(t *testing.T, expectation, result string) {
		t.Helper()
		if result != expectation {
			t.Errorf("result '%s', expected '%s'", result, expectation)
		}
	}

	t.Run("Say hello to people", func(t *testing.T) {
		result := Hello("Test")
		expectation := "Hello Test"

		validateMessage(t, expectation, result)
	})

	t.Run("If empty string is provided, should return hello world", func(t *testing.T) {
		result := Hello("")
		expectation := "Hello World"

		validateMessage(t, expectation, result)
	})
}
