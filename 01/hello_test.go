package main

import "testing"

func TestHello(t *testing.T) {
	validateMessage := func(t *testing.T, expectation, result string) {
		t.Helper()
		if result != expectation {
			t.Errorf("result '%s', expected '%s'", result, expectation)
		}
	}

	t.Run("Say hello to people", func(t *testing.T) {
		result := Hello("Teste")
		expectation := "Ola Teste"

		validateMessage(t, expectation, result)
	})

	t.Run("If empty string is provided, should return hello world", func(t *testing.T) {
		result := Hello("")
		expectation := "Ola Mundo"

		validateMessage(t, expectation, result)
	})
}
